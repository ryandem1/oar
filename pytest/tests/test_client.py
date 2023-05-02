import typing

import pytest
import requests_mock

import oar


class TestClient:

    def test_add_result(self, valid_test: dict[str, typing.Any]):
        """
        Test that adding a result works properly.

        Parameters
        ----------
        valid_test : dict[str, typing.Any]
            Valid OAR test
        """
        client = oar.Client("http://localhost:8080")
        test = oar.Test(**valid_test)

        with requests_mock.Mocker() as m:
            m.post("http://localhost:8080/test", json=test.id_)
            test_id = client.add_test(oar.Test(**valid_test))

        assert test_id == test.id_

    def test_add_result_silently_fails(self, valid_test: dict[str, typing.Any]):
        """
        Test that failure to add a result because of a server failure doesn't result in an exception.

        Parameters
        ----------
        valid_test : dict[str, typing.Any]
            Valid OAR test
        """
        client = oar.Client("http://fake-bad-address:8080")
        assert not client.add_test(oar.Test(**valid_test))

    def test_query(self, valid_query: dict[str, typing.Any]):
        """
        Test that the /query handler works

        Parameters
        ----------
        valid_query: dict[str, typing.Any]
            Query
        """
        query = oar.TestQuery(**valid_query)
        client = oar.Client("http://localhost:8080")

        with requests_mock.Mocker() as m:
            m.post("http://localhost:8080/query", json=query.as_query_string())
            query_string = client.query(query)

        assert oar.TestQuery.from_query_string(query_string) == query

    def test_query_soft_fails(self, valid_query: dict[str, typing.Any]):
        """
        Test that the /query handler fails softly

        Parameters
        ----------
        valid_query: dict[str, typing.Any]
            Query
        """
        query = oar.TestQuery(**valid_query)
        client = oar.Client("http://localhost:8080")

        with requests_mock.Mocker() as m:
            m.post("http://localhost:8080/query", status_code=400)
            query_string = client.query(query)

        assert not query_string

    def test_get_tests(self, valid_query: dict[str, typing.Any], valid_test: dict[str, typing.Any]):
        """
        Test that the GET /tests handler works

        Parameters
        ----------
        valid_query: dict[str, typing.Any]
            Query

        valid_test: dict[str, typing.Any]
            Test
        """
        query = oar.TestQuery(**valid_query)
        client = oar.Client("http://localhost:8080")
        query_result = oar.TestQueryResult(
            count=1,
            tests=[oar.Test(**valid_test)]
        )

        with requests_mock.Mocker() as m:
            m.get(
                url="http://localhost:8080/tests",
                json=query_result.dict()
            )
            results = client.get_tests(query)

        assert results == query_result

    def test_get_tests_soft_fail(self, valid_query: dict[str, typing.Any], valid_test: dict[str, typing.Any]):
        """
        Test that the GET /tests handler fails softly

        Parameters
        ----------
        valid_query: dict[str, typing.Any]
            Query

        valid_test: dict[str, typing.Any]
            Test
        """
        query = oar.TestQuery(**valid_query)
        client = oar.Client("http://localhost:8080")

        with requests_mock.Mocker() as m:
            m.get(
                url="http://localhost:8080/tests",
                status_code=400
            )
            results = client.get_tests(query)

        assert not results

    @pytest.mark.parametrize("json_data", [True, False])
    def test_invalid_response_logging(self, valid_query: dict[str, typing.Any], json_data):
        """
        Test that logging paths work correctly for invalid responses

        Parameters
        ----------
        valid_query: dict[str, typing.Any]
            Test
        """
        query = oar.TestQuery(**valid_query)
        client = oar.Client("http://localhost:8080")

        with requests_mock.Mocker() as m:
            m.get(
                url="http://localhost:8080/tests",
                status_code=400,
                json={"error": "some fake error"} if json_data else None,
                text="{invalid}" if not json_data else None
            )
            client.get_tests(query)

    def test_enrich_tests(self, valid_query: dict[str, typing.Any], valid_test: dict[str, typing.Any]):
        """
        Test that the PATCH /tests handler works

        Parameters
        ----------
        valid_query: dict[str, typing.Any]
            Query

        valid_test: dict[str, typing.Any]
            Test
        """
        query = oar.TestQuery(**valid_query)
        test = oar.Test(**valid_test)
        client = oar.Client("http://localhost:8080")

        with requests_mock.Mocker() as m:
            m.patch(
                url="http://localhost:8080/tests",
                status_code=200
            )
            response_code = client.enrich_tests(test, query)

        assert response_code

    def test_delete_tests(self, valid_query: dict[str, typing.Any]):
        """
        Test that the DELETE /tests handler works

        Parameters
        ----------
        valid_query: dict[str, typing.Any]
            Query
        """
        query = oar.TestQuery(**valid_query)
        client = oar.Client("http://localhost:8080")

        with requests_mock.Mocker() as m:
            m.delete(
                url="http://localhost:8080/tests",
                status_code=200
            )
            response_code = client.delete_tests(query)

        assert response_code
