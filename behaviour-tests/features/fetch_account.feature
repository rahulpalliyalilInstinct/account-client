Feature: fetch my account
  In order to fetch my account
  As a user
  I need to call the fetch account api

  Scenario Outline: Send request to fetch account
    Given  a new registered account "<accountid>" with countryCode "<countryCode>"
    Then I am able to see my account  "<accountid>" details

    Examples:
      | accountid |                                 |countryCode|
      | ad27e265-9605-4b4b-a0e5-3005ea8cc1df |      |    GB        |