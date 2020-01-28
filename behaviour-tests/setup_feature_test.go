package behaviour

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
)

func FeatureContext(s *godog.Suite) {
	var testRegisterClient = NewRegisterTestClient()
	var testDeRegisterClient = NewTestDeRegisterClient()
	var testListClient = NewListTestClient()
	var testFetchClient = NewFetchTestClient()
	fn := func(feature *gherkin.Feature) {
		testRegisterClient.cleanUp()
		testListClient.cleanUp()
		testFetchClient.cleanUp()
	}

	s.AfterFeature(fn)
	s.Step(`^a new unregistered account "([^"]*)" with countryCode "([^"]*)"$`, testRegisterClient.aNewUnregisteredAccountWithCountryCode)
	s.Step(`^I send a request to register the account$`, testRegisterClient.iSendARequestToRegisterTheAccount)
	s.Step(`^I am able to see my account "([^"]*)" registered$`, testRegisterClient.iAmAbleToSeeMyAccountRegistered)

	s.Step(`^a new registered account "([^"]*)" with countryCode "([^"]*)"$`, testDeRegisterClient.aNewRegisteredAccountWithCountryCode)
	s.Step(`^I send a request to deregister the account$`, testDeRegisterClient.iSendARequestToDeregisterTheAccount)
	s.Step(`^I am able to see my account "([^"]*)" deregistered$`, testDeRegisterClient.iAmAbleToSeeMyAccountDeregistered)

	s.Step(`^multiple accounts with "([^"]*)" and "([^"]*)" within same country "([^"]*)"$`, testListClient.multipleAccountsWithAndWithinSameCountry)
	s.Step(`^I want to list a single account per page$`, testListClient.iWantToListASingleAccountPerPage)
	s.Step(`^I am able to see my first account in page (\d+) and second account in page (\d+)$`, testListClient.iAmAbleToSeeMyFirstAccountInPageAndSecondAccountInPage)

	s.Step(`^I am able to see my account  "([^"]*)" details$`, testFetchClient.iAmAbleToSeeMyAccountDetails)
}
