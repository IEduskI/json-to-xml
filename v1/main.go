package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

// Helper function to convert map to XML
func mapToXML(data map[string]interface{}) ([]byte, error) {
	return xml.MarshalIndent(&MapWrapper{Items: data}, "", "  ")
}

type MapWrapper struct {
	XMLName xml.Name
	Items   map[string]interface{}
}

func (mw *MapWrapper) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if mw.XMLName.Local == "" {
		start.Name.Local = "root"
	} else {
		start.Name = mw.XMLName
	}
	tokens := []xml.Token{start}

	for key, value := range mw.Items {
		elem := xml.StartElement{Name: xml.Name{Local: key}}
		tokens = append(tokens, elem)

		switch v := value.(type) {
		case map[string]interface{}:
			tokens = append(tokens, mapToTokens(v)...)
		case []interface{}:
			for _, item := range v {
				tokens = append(tokens, arrayToTokens(key, item)...)
			}
		default:
			if v != nil {
				tokens = append(tokens, xml.CharData(fmt.Sprintf("%v", v)))
			} else {
				tokens = append(tokens, xml.CharData(fmt.Sprintf("%v", "")))
			}
		}

		tokens = append(tokens, xml.EndElement{Name: elem.Name})
	}

	tokens = append(tokens, xml.EndElement{Name: start.Name})
	for _, t := range tokens {
		if err := e.EncodeToken(t); err != nil {
			return err
		}
	}
	return e.Flush()
}

func mapToTokens(data map[string]interface{}) []xml.Token {
	var tokens []xml.Token
	for key, value := range data {
		elem := xml.StartElement{Name: xml.Name{Local: key}}
		tokens = append(tokens, elem)

		switch v := value.(type) {
		case map[string]interface{}:
			tokens = append(tokens, mapToTokens(v)...)
		case []interface{}:
			for _, item := range v {
				tokens = append(tokens, arrayToTokens(key, item)...)
			}
		default:
			if v != nil {
				tokens = append(tokens, xml.CharData(fmt.Sprintf("%v", v)))
			} else {
				tokens = append(tokens, xml.CharData(fmt.Sprintf("%v", "")))
			}
		}

		tokens = append(tokens, xml.EndElement{Name: elem.Name})
	}
	return tokens
}

func arrayToTokens(key string, item interface{}) []xml.Token {
	var tokens []xml.Token
	elem := xml.StartElement{Name: xml.Name{Local: key}}
	tokens = append(tokens, elem)

	switch v := item.(type) {
	case map[string]interface{}:
		tokens = append(tokens, mapToTokens(v)...)
	case []interface{}:
		for _, subItem := range v {
			tokens = append(tokens, arrayToTokens(key, subItem)...)
		}
	default:
		if v != nil {
			tokens = append(tokens, xml.CharData(fmt.Sprintf("%v", v)))
		} else {
			tokens = append(tokens, xml.CharData(fmt.Sprintf("%v", "")))
		}
	}

	tokens = append(tokens, xml.EndElement{Name: elem.Name})
	return tokens
}

func main() {
	// Sample JSON data
	jsonData := `{
		"companyCode": "LM",
		"programCode": "LMS",
		"membershipNumber": "63534237605",
		"activityCode": "BP",
		"txnHeader": {
			"transactionID": "7EDB4501BF04CCFB8797AFAABC146AB176C73CDE",
			"userName": "IBSADMIN",
			"channelUserCode": "IBSADMIN-63534237605",
			"transactionToken": "",
			"timeStamp": "2023-10-16T22:50:52.019Z",
			"deviceId": "",
			"deviceIP": "192.168.250.210",
			"deviceOperatingSystem": "",
			"deviceLocationLatitude": "",
			"deviceLocationLongitude": "",
			"deviceCountryCode": "",
			"additionalInfo": "",
			"remarks": ""
		},
		"acceptPayment": {
			"activityType": "BP",
			"paymentRemark": null,
			"promotionIdentifier": null,
			"chargeHeadDetail": null,
			"waiverDetail": null,
			"paymentDetail": [
				{
					"pointsCollected": 0.0,
					"amountCollected": 129.79,
					"paymentType": "CC",
					"externalProgramCode": null,
					"currencyCode": "USD",
					"paymentSource": "CTO",
					"quoteReferenceNumber": null,
					"paymentGateWayRefNumber": null,
					"bankRefNumber": null,
					"cardType": null,
					"cardHolderName": "SALINAS\\ISMAEL",
					"cardNumber": null,
					"expiryDate": null
				},
				{
					"pointsCollected": 0.0,
					"amountCollected": 129.79,
					"paymentType": "CC",
					"externalProgramCode": null,
					"currencyCode": "USD",
					"paymentSource": "CTO",
					"quoteReferenceNumber": null,
					"paymentGateWayRefNumber": null,
					"bankRefNumber": null,
					"cardType": null,
					"cardHolderName": "SALINAS\\ISMAEL",
					"cardNumber": null,
					"expiryDate": null
				}
			],
			"paymentDocuments": [
				{
					"documentNumber": "2024080248639",
					"documentType": "S",
					"fullprintURL": null,
					"printingID": null,
					"USDValue": "129.79",
					"USDTaxValue": "0.0",
					"points": "4000.0",
					"serviceType": "DM",
					"transactionType": "FLX",
					"issueDate": null,
					"transactionValue": "129.79",
					"transactionTaxValue": "0.0",
					"exchangeRate": "1.0",
					"currencyCode": "USD",
					"taxExempt": "false",
					"taxCode": "SV",
					"taxPercent": "0",
					"taxIdentifier": "",
					"issueUser": "IBSADMIN",
					"formOfPayment": "CC",
					"costCenter": null,
					"authUser": null,
					"documentDynamicAttribute": [
						{
							"attributeKey": "COUNTRY",
							"attributeValue": "SV"
						},
						{
							"attributeKey": "STATION",
							"attributeValue": "SAL"
						},
						{
							"attributeKey": "IATA",
							"attributeValue": "44123456"
						}
					]
				},
				{
					"documentNumber": "2024080248639",
					"documentType": "S",
					"fullprintURL": null,
					"printingID": null,
					"USDValue": "129.79",
					"USDTaxValue": "0.0",
					"points": "4000.0",
					"serviceType": "DM",
					"authUser": null,
					"documentDynamicAttribute": [
						{
							"attributeKey": "COUNTRY",
							"attributeValue": "SV"
						},
						{
							"attributeKey": "STATION",
							"attributeValue": "SAL"
						},
						{
							"attributeKey": "IATA",
							"attributeValue": "44123456"
						}
					]
				}
			]
		}
	}`

	// Unmarshal the JSON into a map
	var resultMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &resultMap)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Convert the map to XML
	xmlData, err := mapToXML(resultMap)
	if err != nil {
		log.Fatalf("Error converting map to XML: %v", err)
	}

	// Print the XML result
	fmt.Println(string(xmlData))
}
