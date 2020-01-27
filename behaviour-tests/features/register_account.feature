Feature: registering an account
  In order to register an account
  As a user
  I need to call the create account api

  Scenario Outline: Send request to register an account
    Given a new unregistered account "<accountid>" with countryCode "<countryCode>"
    When I send a request to register the account
    Then I am able to see my account "<accountid>" registered

    Examples:
      | accountid |                                 |countryCode|
      | ad27e265-9605-4b4b-a0e5-3005ea8cc1da |      |    GB        |
