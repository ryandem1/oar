import typing

import oar


class TestReport:

    def test_add_test_to_report(self, valid_test: dict[str, typing.Any]):
        """
        Tests adding a test to a report works
        """
        report = oar.Report()
        test = oar.Test(**valid_test)
        report += test

        assert len(report.tests) == 1
        assert report.tests[0] == test
