import typing

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
