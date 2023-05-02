"""
Unit tests for the ``models.py`` module. "Tests" were called "results" here for some clarity.
"""
import base64
import enum
import typing

import oar


class TestModels:

    def test_valid_result(self, valid_test: dict[str, typing.Any]):
        """
        Ensures a valid test result is accepted.

        Parameters
        ----------
        valid_test : dict[str, typing.Any]
            Valid JSON test
        """
        oar.Test(**valid_test)

    def test_as_request_body(self, valid_test: dict[str, typing.Any]):
        """
        Ensures request body is properly formatted.

        Parameters
        ----------
        valid_test : dict[str, typing.Any]
            Valid JSON test
        """
        test = oar.Test(**valid_test)
        body = test.as_request_body()

        for v in body.values():
            assert not isinstance(v, enum.Enum)

    def test_query_init(self, valid_query: dict[str, typing.Any]):
        """
        Test that valid queries initialize correctly.

        Parameters
        ----------
        valid_query : dict[str, typing.Any]
            Valid test query
        """
        oar.TestQuery(**valid_query)

    def test_query_can_encode(self, valid_query: dict[str, typing.Any]):
        """
        Tests that a query can encode into valid base64.

        Parameters
        ----------
        valid_query : dict[str, typing.Any]
            Valid test query
        """
        query = oar.TestQuery(**valid_query)
        query_string = query.as_query_string()

        base64.b64decode(query_string)  # Should decode successfully

    def test_query_can_decode(self, valid_query: dict[str, typing.Any]):
        """
        Tests that a query can decode back into another query

        Parameters
        ----------
        valid_query : dict[str, typing.Any]
            Valid test query
        """
        query = oar.TestQuery(**valid_query)
        query_string = query.as_query_string()

        decoded_query = oar.TestQuery.from_query_string(query_string)
        assert query == decoded_query

    def test_query_as_request_body(self, valid_query: dict[str, typing.Any]):
        """
        Ensures a valid query can be a request body.

        Parameters
        ----------
        valid_query : dict[str, typing.Any]
            Valid test query
        """
        query = oar.TestQuery(**valid_query)
        request_body = query.as_request_body()
        request_body_query = oar.TestQuery(**request_body)  # Should go both ways

        assert query == request_body_query
