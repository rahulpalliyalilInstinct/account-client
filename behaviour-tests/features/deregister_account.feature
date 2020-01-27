Feature: deregistering an account
  In order to deregister an account
  As a user
  I need to call the delete account api

  Scenario Outline: Send request to deregister an account
    Given a new registered account "<accountid>" with countryCode "<countryCode>"
    When I send a request to deregister the account
    Then I am able to see my account "<accountid>" deregistered

    Examples:
      | accountid |                                 |countryCode|
      | ad27e265-9605-4b4b-a0e5-3005ea8cc1da |      |    GB        |