Feature: list my accounts
  In order to list my accounts based on page number and size
  As a user
  I need to call the list accounts api

  Scenario Outline: Send request to list accounts
    Given multiple accounts with "<accountid1>" and "<accountid2>" within same country "<countryCode>"
    When I want to list a single account per page
    Then I am able to see my first account in page 1 and second account in page 2

    Examples:
      | accountid1 |                           |accountid2|                                 |countryCode|
      | ad27e265-9605-4b4b-a0e5-3005ea8cc1da | | ad27e265-9605-4b4b-a0e5-3005ea8cc1db |    |    GB        |


