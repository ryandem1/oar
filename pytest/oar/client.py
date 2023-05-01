import json
from logging import getLogger

from requests import Session, Response
from requests.adapters import Retry, HTTPAdapter
from requests.exceptions import RetryError, ConnectionError

from oar.models import Test, TestQuery, TestQueryResult

logger = getLogger("oar")


class Client:
    """
    Client that provides a Python interface over an OAR HTTP client
    """

    def __init__(self, base_url: str, session: Session = Session()):
        """
        Initializes the client with a ``base_url`` for OAR, as well as a Session

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
        self.query_route = self.base_url + "/query"

        retries = Retry(total=4, allowed_methods=False, backoff_factor=0.5, status_forcelist=[400, 500, 502, 503, 504])
        self.session.mount('http://', HTTPAdapter(max_retries=retries))
        self.session.mount('https://', HTTPAdapter(max_retries=retries))

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
                message = None
                try:
                    message = response.json()
                except json.JSONDecodeError:
                    error_message += f"\nStatus Code: {response.status_code}\nText: {response.text}"

                error_message += f"\nStatus Code: {response.status_code}\nMessage: {message}"
            logger.error(error_message)

    def add_test(self, test: Test) -> int | None:
        """
        Sends a POST to the ``/test`` endpoint to add a new test result.

        Parameters
        ----------
        test : Test
            OAR test to add

        Returns
        -------
        test_id : int | None
            ID of the created test. Will return None on error
        """
        try:
            response = self.session.post(self.test_route, json=test.as_request_body())
        except (RetryError, ConnectionError):
            logger.error("Max retries reached for adding test result, continuing but you should probably look at this!")
            return  # Will silently fail

        self.__log_error_if_not_ok(response)
        test_id = response.json() if response.ok else None
        return test_id

    def query(self, query: TestQuery) -> str:
        """
        Will call the ``/query`` endpoint with a TestQuery and return a base64 encoded string to use as the "query"
        parameter for the ``/tests`` methods.

        Notes
        -----
        This is here for a complete interface, but you will not need to use this if just working in Python.
        ``TestQuery`` already provides a ``.as_query_string()`` method that can be used without needing to call the
        endpoint.

        Parameters
        ----------
        query : TestQuery
            Query to encode

        Returns
        -------
        query_string : str
            base64 encoded string returned from the endpoint
        """
        response = self.session.post(self.query_route, json=query.as_request_body())
        self.__log_error_if_not_ok(response)
        return response.json()

    def get_tests(self, query: TestQuery, offset: int = 0, limit: int = 250) -> TestQueryResult | None:
        """
        Sends a GET to the ``/tests`` endpoint to return the results of the test query.

        Parameters
        ----------
        query : TestQuery
            Query to get the results of

        offset : int
            Query offset

        limit : int
            Query limit

        Returns
        -------
        result : TestQueryResult | None
            Result of the query. Will return None if error
        """
        response = self.session.get(
            url=self.tests_route,
            params={
                "offset": offset,
                "limit": limit,
                "query": query.as_query_string()
            }
        )
        self.__log_error_if_not_ok(response)
        result = TestQueryResult(**response.json()) if response.ok else None
        return result

    def enrich_tests(self, test: Test, query: TestQuery) -> int:
        """
        Sends a PATCH to the ``/tests`` endpoint to enrich all tests that match the query passed

        Parameters
        ----------
        test : Test
            Test details to enrich existing result with

        query : TestQuery
            Will enrich the results of the test query passed.

        Returns
        -------
        status_code : int
            Status code which indicates: 304 if no tests were found with those IDs or the request failed, or else will
            return a 200 if tests were modified.
        """
        response = self.session.patch(
            url=self.tests_route,
            json=test.as_request_body(),
            params={"query": query.as_query_string()}
        )
        self.__log_error_if_not_ok(response)
        return response.status_code

    def delete_tests(self, query: TestQuery) -> int:
        """
        Will send a DELETE to the ``/tests`` endpoint to delete tests by test query.

        Parameters
        ----------
        query : TestQuery
            Test query to delete the results of

        Returns
        -------
        status_code : int
            Status code which indicates: 304 if no tests were found with those IDs or the request failed, or else will
            return a 200 if tests were deleted.
        """
        response = self.session.delete(
            url=self.tests_route,
            params={"query": query.as_query_string()}
        )
        self.__log_error_if_not_ok(response)
        return response.status_code
