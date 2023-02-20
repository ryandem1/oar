from requests import Session, Response
from oar.models import Test
from logging import getLogger

logger = getLogger(__name__)


class Client:
    """
    Client that provides a Python interface over an OAR HTTP client
    """

    def __init__(self, base_url: str, session: Session = Session()):
        """
        Initializes the client with a ``base_url`` for OAR, as well as a Session.

        Parameters
        ----------
        base_url : str
            Base URL of the OAR instance

        session : Session
            Requests session for the client to use. By default, will make its own
        """
        self.base_url = base_url
        self.session = session
        self.test_route = self.base_url + "/test"
        self.tests_route = self.base_url + "/tests"

    @staticmethod
    def __log_error_if_not_ok(response: Response) -> None:
        """
        Will log out an error if the response return is not of 2xx status. This is designed to not stop tests if things
        go wrong so that if it is used in a test, it will not cause false positives.

        Parameters
        ----------
        response : Response
            Requests response object to check

        Returns
        -------
        None
        """
        if not response.ok:
            error_message = "Error adding OAR test! Continuing, but you should probably look at this!"
            if response.text:
                error_message += f"\nMessage: {response.json()}"
            logger.error(error_message)

    def add_test(self, test: Test) -> int:
        """
        Sends a POST to the ``/test`` endpoint to add a new test result.

        Parameters
        ----------
        test : Test
            OAR test to add

        Returns
        -------
        test_id : int
            ID of the created test
        """
        response = self.session.post(self.test_route, json=test.as_request_body())
        self.__log_error_if_not_ok(response)
        test_id = response.json()
        return test_id

    def enrich_test(self, test: Test) -> None:
        """
        Sends a PATCH to the ``/test`` endpoint to enrich an existing test result.

        Parameters
        ----------
        test : Test
            Test details to enrich existing result with

        Returns
        -------
        None
        """
        response = self.session.patch(self.test_route, json=test.as_request_body())
        self.__log_error_if_not_ok(response)

    def delete_tests(self, *test_ids: int) -> int:
        """
        Will send a DELETE to the ``/tests`` endpoint to delete tests by IDs. Will return the status code of the request

        Parameters
        ----------
        test_ids : int
            IDs of the tests to be deleted.

        Returns
        -------
        status_code : int
            Status code which indicates: 304 if no tests were found with those IDs or the request failed, or else will
            return a 200 if tests were deleted.
        """
        body = [{"ID": id_} for id_ in test_ids]
        response = self.session.delete(self.tests_route, json=body)
        self.__log_error_if_not_ok(response)
        return 200 if response.status_code == 200 else 304
