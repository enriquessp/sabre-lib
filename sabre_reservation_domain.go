package sabre

import (
	"encoding/xml"

	"github.com/agentguru/collections-lib"
)

//  REQUEST

type GetReservationRQ struct {
	XMLName       xml.Name      `xml:"ns4:GetReservationRQ"`
	Version       string        `xml:"Version,attr"`
	Ns2           string        `xml:"xmlns:ns2,attr"`
	Ns4           string        `xml:"xmlns:ns4,attr"`
	Ns5           string        `xml:"xmlns:ns5,attr"`
	Ns6           string        `xml:"xmlns:ns6,attr"`
	Ns7           string        `xml:"xmlns:ns7,attr"`
	Ns8           string        `xml:"xmlns:ns8,attr"`
	Locator       string        `xml:"ns4:Locator"`
	RequestType   string        `xml:"ns4:RequestType"`
	ReturnOptions ReturnOptions `xml:"ns4:ReturnOptions"`
}

type ReturnOptions struct {
	PriceQuoteServiceVersion string       `xml:"PriceQuoteServiceVersion,attr"`
	Xsi                      string       `xml:"xmlns:xsi,attr"`
	Type                     string       `xml:"xsi:type,attr"`
	SubjectAreas             SubjectAreas `xml:"ns4:SubjectAreas"`
	ViewName                 string       `xml:"ns4:ViewName"`
	ResponseFormat           string       `xml:"ns4:ResponseFormat"`
}

type SubjectAreas struct {
	SubjectArea []string `xml:"ns4:SubjectArea"`
}

// REQUEST

type GetReservationRS struct {
	Text        string      `xml:",chardata"`
	Ns4         string      `xml:"ns4,attr"`
	Ns6         string      `xml:"ns6,attr"`
	Or114       string      `xml:"or114,attr"`
	Raw         string      `xml:"raw,attr"`
	Stl19       string      `xml:"stl19,attr"`
	Version     string      `xml:"Version,attr"`
	Reservation Reservation `xml:"Reservation"`
	PriceQuote  struct {
		Text           string `xml:",chardata"`
		PriceQuoteInfo struct {
			Text        string `xml:",chardata"`
			Xmlns       string `xml:"xmlns,attr"`
			Reservation struct {
				Text        string `xml:",chardata"`
				UpdateToken string `xml:"updateToken,attr"`
			} `xml:"Reservation"`
			Summary struct {
				Text            string `xml:",chardata"`
				NameAssociation struct {
					Text       string `xml:",chardata"`
					FirstName  string `xml:"firstName,attr"`
					LastName   string `xml:"lastName,attr"`
					NameId     string `xml:"nameId,attr"`
					NameNumber string `xml:"nameNumber,attr"`
					PriceQuote struct {
						Text         string `xml:",chardata"`
						LatestPQFlag string `xml:"latestPQFlag,attr"`
						Number       string `xml:"number,attr"`
						PricingType  string `xml:"pricingType,attr"`
						Status       string `xml:"status,attr"`
						Type         string `xml:"type,attr"`
						Indicators   string `xml:"Indicators"`
						Passenger    struct {
							Text               string `xml:",chardata"`
							PassengerTypeCount string `xml:"passengerTypeCount,attr"`
							RequestedType      string `xml:"requestedType,attr"`
							Type               string `xml:"type,attr"`
						} `xml:"Passenger"`
						ItineraryType     string `xml:"ItineraryType"`
						ValidatingCarrier string `xml:"ValidatingCarrier"`
						Amounts           struct {
							Text  string `xml:",chardata"`
							Total struct {
								Text         string `xml:",chardata"`
								CurrencyCode string `xml:"currencyCode,attr"`
								DecimalPlace string `xml:"decimalPlace,attr"`
							} `xml:"Total"`
						} `xml:"Amounts"`
						LocalCreateDateTime string `xml:"LocalCreateDateTime"`
					} `xml:"PriceQuote"`
				} `xml:"NameAssociation"`
			} `xml:"Summary"`
			Details []struct {
				Text          string `xml:",chardata"`
				Number        string `xml:"number,attr"`
				PassengerType string `xml:"passengerType,attr"`
				PricingType   string `xml:"pricingType,attr"`
				Status        string `xml:"status,attr"`
				Type          string `xml:"type,attr"`
				AgentInfo     struct {
					Text         string `xml:",chardata"`
					Duty         string `xml:"duty,attr"`
					Sine         string `xml:"sine,attr"`
					HomeLocation string `xml:"HomeLocation"`
					WorkLocation string `xml:"WorkLocation"`
				} `xml:"AgentInfo"`
				TransactionInfo struct {
					Text                string `xml:",chardata"`
					CreateDateTime      string `xml:"CreateDateTime"`
					UpdateDateTime      string `xml:"UpdateDateTime"`
					LastDateToPurchase  string `xml:"LastDateToPurchase"`
					LocalCreateDateTime string `xml:"LocalCreateDateTime"`
					LocalUpdateDateTime string `xml:"LocalUpdateDateTime"`
					InputEntry          string `xml:"InputEntry"`
				} `xml:"TransactionInfo"`
				NameAssociationInfo struct {
					Text       string `xml:",chardata"`
					FirstName  string `xml:"firstName,attr"`
					LastName   string `xml:"lastName,attr"`
					NameId     string `xml:"nameId,attr"`
					NameNumber string `xml:"nameNumber,attr"`
				} `xml:"NameAssociationInfo"`
				SegmentInfo []struct {
					Text          string `xml:",chardata"`
					Number        string `xml:"number,attr"`
					SegmentStatus string `xml:"segmentStatus,attr"`
					Flight        struct {
						Text                string `xml:",chardata"`
						ConnectionIndicator string `xml:"connectionIndicator,attr"`
						MarketingFlight     struct {
							Text   string `xml:",chardata"`
							Number string `xml:"number,attr"`
						} `xml:"MarketingFlight"`
						ClassOfService string `xml:"ClassOfService"`
						Departure      struct {
							Text     string `xml:",chardata"`
							DateTime string `xml:"DateTime"`
							CityCode struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
							} `xml:"CityCode"`
						} `xml:"Departure"`
						Arrival struct {
							Text     string `xml:",chardata"`
							DateTime string `xml:"DateTime"`
							CityCode struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
							} `xml:"CityCode"`
						} `xml:"Arrival"`
					} `xml:"Flight"`
					FareBasis      string `xml:"FareBasis"`
					NotValidBefore string `xml:"NotValidBefore"`
					NotValidAfter  string `xml:"NotValidAfter"`
					Baggage        struct {
						Text      string `xml:",chardata"`
						Allowance string `xml:"allowance,attr"`
						Type      string `xml:"type,attr"`
					} `xml:"Baggage"`
					BrandedFare string `xml:"BrandedFare"`
				} `xml:"SegmentInfo"`
				FareInfo struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					FareIndicators struct {
						Text           string `xml:",chardata"`
						NegotiatedFare string `xml:"negotiatedFare,attr"`
					} `xml:"FareIndicators"`
					BaseFare struct {
						Text         string `xml:",chardata"`
						CurrencyCode string `xml:"currencyCode,attr"`
						DecimalPlace string `xml:"decimalPlace,attr"`
					} `xml:"BaseFare"`
					TotalTax struct {
						Text         string `xml:",chardata"`
						CurrencyCode string `xml:"currencyCode,attr"`
						DecimalPlace string `xml:"decimalPlace,attr"`
					} `xml:"TotalTax"`
					TotalFare struct {
						Text         string `xml:",chardata"`
						CurrencyCode string `xml:"currencyCode,attr"`
						DecimalPlace string `xml:"decimalPlace,attr"`
					} `xml:"TotalFare"`
					TaxInfo struct {
						Text        string `xml:",chardata"`
						CombinedTax []struct {
							Text   string `xml:",chardata"`
							Code   string `xml:"code,attr"`
							Amount struct {
								Text         string `xml:",chardata"`
								CurrencyCode string `xml:"currencyCode,attr"`
								DecimalPlace string `xml:"decimalPlace,attr"`
							} `xml:"Amount"`
						} `xml:"CombinedTax"`
						Tax []struct {
							Text   string `xml:",chardata"`
							Code   string `xml:"code,attr"`
							Amount struct {
								Text         string `xml:",chardata"`
								CurrencyCode string `xml:"currencyCode,attr"`
								DecimalPlace string `xml:"decimalPlace,attr"`
							} `xml:"Amount"`
						} `xml:"Tax"`
					} `xml:"TaxInfo"`
					FareCalculation struct {
						Text           string `xml:",chardata"`
						RateOfExchange string `xml:"rateOfExchange,attr"`
					} `xml:"FareCalculation"`
					FareComponent []struct {
						Text                 string `xml:",chardata"`
						FareBasisCode        string `xml:"fareBasisCode,attr"`
						Number               string `xml:"number,attr"`
						Type                 string `xml:"type,attr"`
						TypeBitmap           string `xml:"typeBitmap,attr"`
						FlightSegmentNumbers struct {
							Text          string   `xml:",chardata"`
							SegmentNumber []string `xml:"SegmentNumber"`
						} `xml:"FlightSegmentNumbers"`
						FareDirectionality struct {
							Text      string `xml:",chardata"`
							RoundTrip string `xml:"roundTrip,attr"`
							Inbound   string `xml:"inbound,attr"`
						} `xml:"FareDirectionality"`
						Departure struct {
							Text     string `xml:",chardata"`
							DateTime string `xml:"DateTime"`
							CityCode struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
							} `xml:"CityCode"`
						} `xml:"Departure"`
						Arrival struct {
							Text     string `xml:",chardata"`
							DateTime string `xml:"DateTime"`
							CityCode struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
							} `xml:"CityCode"`
						} `xml:"Arrival"`
						Amount struct {
							Text         string `xml:",chardata"`
							CurrencyCode string `xml:"currencyCode,attr"`
							DecimalPlace string `xml:"decimalPlace,attr"`
						} `xml:"Amount"`
						GoverningCarrier string `xml:"GoverningCarrier"`
						Rules            struct {
							Text    string `xml:",chardata"`
							Carrier string `xml:"carrier,attr"`
							Rule    string `xml:"rule,attr"`
							Tariff  string `xml:"tariff,attr"`
							Vendor  string `xml:"vendor,attr"`
						} `xml:"Rules"`
					} `xml:"FareComponent"`
				} `xml:"FareInfo"`
				FeeInfo           string `xml:"FeeInfo"`
				MiscellaneousInfo struct {
					Text              string `xml:",chardata"`
					ValidatingCarrier string `xml:"ValidatingCarrier"`
					ItineraryType     string `xml:"ItineraryType"`
				} `xml:"MiscellaneousInfo"`
				MessageInfo struct {
					Text    string `xml:",chardata"`
					Message []struct {
						Text   string `xml:",chardata"`
						Number string `xml:"number,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"Message"`
					Remarks struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"Remarks"`
					PricingParameters string `xml:"PricingParameters"`
					BaggageDisclosure string `xml:"BaggageDisclosure"`
				} `xml:"MessageInfo"`
				HistoryInfo struct {
					Text      string `xml:",chardata"`
					AgentInfo struct {
						Text         string `xml:",chardata"`
						Sine         string `xml:"sine,attr"`
						HomeLocation string `xml:"HomeLocation"`
					} `xml:"AgentInfo"`
					TransactionInfo struct {
						Text          string `xml:",chardata"`
						LocalDateTime string `xml:"LocalDateTime"`
						InputEntry    string `xml:"InputEntry"`
					} `xml:"TransactionInfo"`
				} `xml:"HistoryInfo"`
			} `xml:"Details"`
		} `xml:"PriceQuoteInfo"`
	} `xml:"PriceQuote"`
}

func (r GetReservationRS) GetOpenReservationsByPax(paxId string) []OpenReservationElement {
	return collections.FilterSlice(r.Reservation.OpenReservationElements.OpenReservationElement, func(o OpenReservationElement) bool { return o.NameAssociation.NameRefNumber == paxId })
}

type Reservation struct {
	Text                    string                  `xml:",chardata"`
	NumberInSegment         string                  `xml:"NumberInSegment,attr"`
	NumberInParty           string                  `xml:"numberInParty,attr"`
	NumberOfInfants         string                  `xml:"numberOfInfants,attr"`
	BookingDetails          BookingDetails          `xml:"BookingDetails"`
	POS                     POS                     `xml:"POS"`
	PassengerReservation    PassengerReservation    `xml:"PassengerReservation"`
	ReceivedFrom            ReceivedFrom            `xml:"ReceivedFrom"`
	PhoneNumbers            PhoneNumbers            `xml:"PhoneNumbers"`
	Remarks                 Remarks                 `xml:"Remarks"`
	EmailAddresses          string                  `xml:"EmailAddresses"`
	GenericSpecialRequests  GenericSpecialRequests  `xml:"GenericSpecialRequests"`
	OpenReservationElements OpenReservationElements `xml:"OpenReservationElements"`
}

type BookingDetails struct {
	Text                    string `xml:",chardata"`
	Header                  string `xml:"Header"`
	RecordLocator           string `xml:"RecordLocator"`
	CreationTimestamp       string `xml:"CreationTimestamp"`
	SystemCreationTimestamp string `xml:"SystemCreationTimestamp"`
	CreationAgentID         string `xml:"CreationAgentID"`
	UpdateTimestamp         string `xml:"UpdateTimestamp"`
	PNRSequence             string `xml:"PNRSequence"`
	FlightsRange            struct {
		Text  string `xml:",chardata"`
		End   string `xml:"End,attr"`
		Start string `xml:"Start,attr"`
	} `xml:"FlightsRange"`
	DivideSplitDetails      string `xml:"DivideSplitDetails"`
	EstimatedPurgeTimestamp string `xml:"EstimatedPurgeTimestamp"`
	UpdateToken             string `xml:"UpdateToken"`
}

type POS struct {
	Text        string `xml:",chardata"`
	AirExtras   string `xml:"AirExtras,attr"`
	InhibitCode string `xml:"InhibitCode,attr"`
	Source      struct {
		Text               string `xml:",chardata"`
		AgentDutyCode      string `xml:"AgentDutyCode,attr"`
		AgentSine          string `xml:"AgentSine,attr"`
		AirlineVendorID    string `xml:"AirlineVendorID,attr"`
		BookingSource      string `xml:"BookingSource,attr"`
		HomePseudoCityCode string `xml:"HomePseudoCityCode,attr"`
		ISOCountry         string `xml:"ISOCountry,attr"`
		PrimeHostID        string `xml:"PrimeHostID,attr"`
		PseudoCityCode     string `xml:"PseudoCityCode,attr"`
	} `xml:"Source"`
}

type PassengerReservation struct {
	Text       string `xml:",chardata"`
	Passengers struct {
		Text      string `xml:",chardata"`
		Passenger []struct {
			Text          string `xml:",chardata"`
			ElementId     string `xml:"elementId,attr"`
			ID            string `xml:"id,attr"`
			NameAssocId   string `xml:"nameAssocId,attr"`
			NameId        string `xml:"nameId,attr"`
			NameType      string `xml:"nameType,attr"`
			PassengerType string `xml:"passengerType,attr"`
			LastName      string `xml:"LastName"`
			FirstName     string `xml:"FirstName"`
			FrequentFlyer []struct {
				Text                 string `xml:",chardata"`
				RPH                  string `xml:"RPH,attr"`
				ID                   string `xml:"id,attr"`
				SupplierCode         string `xml:"SupplierCode"`
				Number               string `xml:"Number"`
				TierLevelNumber      string `xml:"TierLevelNumber"`
				ShortText            string `xml:"ShortText"`
				ReceivingCarrierCode string `xml:"ReceivingCarrierCode"`
				StatusCode           string `xml:"StatusCode"`
			} `xml:"FrequentFlyer"`
			SpecialRequests struct {
				Text        string `xml:",chardata"`
				APISRequest struct {
					Text      string `xml:",chardata"`
					DOCSEntry struct {
						Text          string `xml:",chardata"`
						ID            string `xml:"id,attr"`
						Type          string `xml:"type,attr"`
						DateOfBirth   string `xml:"DateOfBirth"`
						Gender        string `xml:"Gender"`
						Surname       string `xml:"Surname"`
						Forename      string `xml:"Forename"`
						MiddleName    string `xml:"MiddleName"`
						PrimaryHolder string `xml:"PrimaryHolder"`
						FreeText      string `xml:"FreeText"`
						ActionCode    string `xml:"ActionCode"`
						NumberInParty string `xml:"NumberInParty"`
						VendorCode    string `xml:"VendorCode"`
					} `xml:"DOCSEntry"`
				} `xml:"APISRequest"`
			} `xml:"SpecialRequests"`
			Seats string `xml:"Seats"`
		} `xml:"Passenger"`
	} `xml:"Passengers"`
	Segments struct {
		Text string `xml:",chardata"`
		Poc  struct {
			Text      string `xml:",chardata"`
			Airport   string `xml:"Airport"`
			Departure string `xml:"Departure"`
		} `xml:"Poc"`
		Segment []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			Sequence string `xml:"sequence,attr"`
			Air      struct {
				Text                        string `xml:",chardata"`
				Code                        string `xml:"Code,attr"`
				CodeShare                   string `xml:"CodeShare,attr"`
				DayOfWeekInd                string `xml:"DayOfWeekInd,attr"`
				ResBookDesigCode            string `xml:"ResBookDesigCode,attr"`
				SmokingAllowed              string `xml:"SmokingAllowed,attr"`
				SpecialMeal                 string `xml:"SpecialMeal,attr"`
				StopQuantity                string `xml:"StopQuantity,attr"`
				ID                          string `xml:"id,attr"`
				IsPast                      string `xml:"isPast,attr"`
				SegmentAssociationId        string `xml:"segmentAssociationId,attr"`
				Sequence                    string `xml:"sequence,attr"`
				DepartureAirport            string `xml:"DepartureAirport"`
				DepartureAirportCodeContext string `xml:"DepartureAirportCodeContext"`
				ArrivalAirport              string `xml:"ArrivalAirport"`
				ArrivalAirportCodeContext   string `xml:"ArrivalAirportCodeContext"`
				ArrivalTerminalName         string `xml:"ArrivalTerminalName"`
				ArrivalTerminalCode         string `xml:"ArrivalTerminalCode"`
				OperatingAirlineCode        string `xml:"OperatingAirlineCode"`
				OperatingAirlineShortName   string `xml:"OperatingAirlineShortName"`
				OperatingFlightNumber       string `xml:"OperatingFlightNumber"`
				EquipmentType               string `xml:"EquipmentType"`
				MarketingAirlineCode        string `xml:"MarketingAirlineCode"`
				MarketingFlightNumber       string `xml:"MarketingFlightNumber"`
				OperatingClassOfService     string `xml:"OperatingClassOfService"`
				MarketingClassOfService     string `xml:"MarketingClassOfService"`
				MarriageGrp                 struct {
					Text     string `xml:",chardata"`
					Ind      string `xml:"Ind"`
					Group    string `xml:"Group"`
					Sequence string `xml:"Sequence"`
				} `xml:"MarriageGrp"`
				Meal []struct {
					Text string `xml:",chardata"`
					Code string `xml:"Code,attr"`
				} `xml:"Meal"`
				Seats                   string `xml:"Seats"`
				AirlineRefId            string `xml:"AirlineRefId"`
				Eticket                 string `xml:"Eticket"`
				DepartureDateTime       string `xml:"DepartureDateTime"`
				ArrivalDateTime         string `xml:"ArrivalDateTime"`
				FlightNumber            string `xml:"FlightNumber"`
				ClassOfService          string `xml:"ClassOfService"`
				ActionCode              string `xml:"ActionCode"`
				NumberInParty           string `xml:"NumberInParty"`
				SegmentSpecialRequests  string `xml:"SegmentSpecialRequests"`
				InboundConnection       string `xml:"inboundConnection"`
				OutboundConnection      string `xml:"outboundConnection"`
				ScheduleChangeIndicator string `xml:"ScheduleChangeIndicator"`
				SegmentBookedDate       string `xml:"SegmentBookedDate"`
				ElapsedTime             string `xml:"ElapsedTime"`
				AirMilesFlown           string `xml:"AirMilesFlown"`
				FunnelFlight            string `xml:"FunnelFlight"`
				ChangeOfGauge           string `xml:"ChangeOfGauge"`
				Cabin                   struct {
					Text      string `xml:",chardata"`
					Code      string `xml:"Code,attr"`
					Lang      string `xml:"Lang,attr"`
					Name      string `xml:"Name,attr"`
					SabreCode string `xml:"SabreCode,attr"`
					ShortName string `xml:"ShortName,attr"`
				} `xml:"Cabin"`
				Banner                string `xml:"Banner"`
				Informational         string `xml:"Informational"`
				DepartureTerminalName string `xml:"DepartureTerminalName"`
				DepartureTerminalCode string `xml:"DepartureTerminalCode"`
			} `xml:"Air"`
			General struct {
				Chardata     string `xml:",chardata"`
				DayOfWeekInd string `xml:"dayOfWeekInd,attr"`
				IsPast       string `xml:"isPast,attr"`
				Line         struct {
					Text   string `xml:",chardata"`
					Number string `xml:"Number,attr"`
					Status string `xml:"Status,attr"`
					Type   string `xml:"Type,attr"`
				} `xml:"Line"`
				Vendor struct {
					Text string `xml:",chardata"`
					Code string `xml:"Code,attr"`
				} `xml:"Vendor"`
				NumberInParty string `xml:"NumberInParty"`
				Location      struct {
					Text         string `xml:",chardata"`
					LocationCode string `xml:"LocationCode,attr"`
				} `xml:"Location"`
				DateTime string `xml:"DateTime"`
				Text     string `xml:"Text"`
			} `xml:"General"`
			Product struct {
				Text           string `xml:",chardata"`
				ID             string `xml:"id,attr"`
				ProductDetails struct {
					Text            string `xml:",chardata"`
					ProductCategory string `xml:"productCategory,attr"`
					ProductName     struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"ProductName"`
					Air struct {
						Text                    string `xml:",chardata"`
						SegmentAssociationId    string `xml:"segmentAssociationId,attr"`
						Sequence                string `xml:"sequence,attr"`
						DepartureAirport        string `xml:"DepartureAirport"`
						ArrivalAirport          string `xml:"ArrivalAirport"`
						ArrivalTerminalName     string `xml:"ArrivalTerminalName"`
						ArrivalTerminalCode     string `xml:"ArrivalTerminalCode"`
						EquipmentType           string `xml:"EquipmentType"`
						MarketingAirlineCode    string `xml:"MarketingAirlineCode"`
						MarketingFlightNumber   string `xml:"MarketingFlightNumber"`
						MarketingClassOfService string `xml:"MarketingClassOfService"`
						MarriageGrp             struct {
							Text     string `xml:",chardata"`
							Group    string `xml:"Group"`
							Sequence string `xml:"Sequence"`
						} `xml:"MarriageGrp"`
						Cabin struct {
							Text      string `xml:",chardata"`
							Code      string `xml:"code,attr"`
							Lang      string `xml:"lang,attr"`
							Name      string `xml:"name,attr"`
							SabreCode string `xml:"sabreCode,attr"`
							ShortName string `xml:"shortName,attr"`
						} `xml:"Cabin"`
						MealCode          []string `xml:"MealCode"`
						ElapsedTime       string   `xml:"ElapsedTime"`
						AirMilesFlown     string   `xml:"AirMilesFlown"`
						FunnelFlight      string   `xml:"FunnelFlight"`
						ChangeOfGauge     string   `xml:"ChangeOfGauge"`
						DisclosureCarrier struct {
							Text   string `xml:",chardata"`
							Code   string `xml:"Code,attr"`
							DOT    string `xml:"DOT,attr"`
							Banner string `xml:"Banner"`
						} `xml:"DisclosureCarrier"`
						AirlineRefId            string `xml:"AirlineRefId"`
						Eticket                 string `xml:"Eticket"`
						DepartureDateTime       string `xml:"DepartureDateTime"`
						ArrivalDateTime         string `xml:"ArrivalDateTime"`
						FlightNumber            string `xml:"FlightNumber"`
						ClassOfService          string `xml:"ClassOfService"`
						ActionCode              string `xml:"ActionCode"`
						NumberInParty           string `xml:"NumberInParty"`
						InboundConnection       string `xml:"inboundConnection"`
						OutboundConnection      string `xml:"outboundConnection"`
						ScheduleChangeIndicator string `xml:"ScheduleChangeIndicator"`
						SegmentBookedDate       string `xml:"SegmentBookedDate"`
						DepartureTerminalName   string `xml:"DepartureTerminalName"`
						DepartureTerminalCode   string `xml:"DepartureTerminalCode"`
					} `xml:"Air"`
				} `xml:"ProductDetails"`
			} `xml:"Product"`
		} `xml:"Segment"`
	} `xml:"Segments"`
	TicketingInfo struct {
		Text               string `xml:",chardata"`
		TicketingTimeLimit struct {
			Text      string `xml:",chardata"`
			ElementId string `xml:"elementId,attr"`
			ID        string `xml:"id,attr"`
			Index     string `xml:"index,attr"`
			Time      string `xml:"Time"`
		} `xml:"TicketingTimeLimit"`
	} `xml:"TicketingInfo"`
	ItineraryPricing struct {
		Text            string `xml:",chardata"`
		PricedItinerary struct {
			Text                    string `xml:",chardata"`
			InputMessage            string `xml:"InputMessage,attr"`
			SequenceNumber          string `xml:"SequenceNumber,attr"`
			StatusCode              string `xml:"StatusCode,attr"`
			TaxExempt               string `xml:"TaxExempt,attr"`
			ValidatingCarrier       string `xml:"ValidatingCarrier,attr"`
			AirItineraryPricingInfo struct {
				Text          string `xml:",chardata"`
				ItinTotalFare struct {
					Text string `xml:",chardata"`
					Code string `xml:"code,attr"`
					Base struct {
						Text         string `xml:",chardata"`
						CurrencyCode string `xml:"currencyCode,attr"`
					} `xml:"Base"`
					Total struct {
						Text         string `xml:",chardata"`
						CurrencyCode string `xml:"currencyCode,attr"`
					} `xml:"Total"`
					Taxes struct {
						Text string `xml:",chardata"`
						Tax  struct {
							Text   string `xml:",chardata"`
							Code   string `xml:"code,attr"`
							Amount struct {
								Text         string `xml:",chardata"`
								Currency     string `xml:"currency,attr"`
								CurrencyCode string `xml:"currencyCode,attr"`
							} `xml:"Amount"`
						} `xml:"Tax"`
						TaxBreakdownCode []struct {
							Text    string `xml:",chardata"`
							TaxPaid string `xml:"TaxPaid,attr"`
						} `xml:"TaxBreakdownCode"`
					} `xml:"Taxes"`
					Totals struct {
						Text string `xml:",chardata"`
						Code string `xml:"code,attr"`
						Base struct {
							Text         string `xml:",chardata"`
							Currency     string `xml:"currency,attr"`
							CurrencyCode string `xml:"currencyCode,attr"`
						} `xml:"Base"`
						TotalTax struct {
							Text         string `xml:",chardata"`
							Currency     string `xml:"currency,attr"`
							CurrencyCode string `xml:"currencyCode,attr"`
						} `xml:"TotalTax"`
						Total struct {
							Text         string `xml:",chardata"`
							Currency     string `xml:"currency,attr"`
							CurrencyCode string `xml:"currencyCode,attr"`
						} `xml:"Total"`
					} `xml:"Totals"`
				} `xml:"ItinTotalFare"`
				PTCFareBreakdown struct {
					Text                  string `xml:",chardata"`
					PassengerTypeQuantity struct {
						Text     string `xml:",chardata"`
						Code     string `xml:"Code,attr"`
						Quantity string `xml:"Quantity,attr"`
					} `xml:"PassengerTypeQuantity"`
					FareBasisCode []string `xml:"FareBasisCode"`
					FareCalc      string   `xml:"FareCalc"`
					FlightSegment []struct {
						Text              string `xml:",chardata"`
						RPH               string `xml:"RPH,attr"`
						DepartureDateTime string `xml:"DepartureDateTime"`
						ResBookDesigCode  string `xml:"ResBookDesigCode"`
						FlightNumber      string `xml:"FlightNumber"`
						ActionCode        string `xml:"ActionCode"`
						FlightType        string `xml:"FlightType"`
						AirPort           struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
						} `xml:"AirPort"`
						OperatingAirline string `xml:"OperatingAirline"`
						FareBasisCode    string `xml:"FareBasisCode"`
						ValidityDates    struct {
							Text           string `xml:",chardata"`
							NotValidBefore string `xml:"NotValidBefore"`
							NotValidAfter  string `xml:"NotValidAfter"`
						} `xml:"ValidityDates"`
					} `xml:"FlightSegment"`
					FareComponent []struct {
						Text                string `xml:",chardata"`
						Amount              string `xml:"Amount,attr"`
						FareBasisCode       string `xml:"FareBasisCode,attr"`
						FareComponentNumber string `xml:"FareComponentNumber,attr"`
						FareDirectionality  string `xml:"FareDirectionality,attr"`
						GoverningCarrier    string `xml:"GoverningCarrier,attr"`
						TicketDesignator    string `xml:"TicketDesignator,attr"`
						Location            struct {
							Text        string `xml:",chardata"`
							Destination string `xml:"Destination,attr"`
							Origin      string `xml:"Origin,attr"`
						} `xml:"Location"`
						Dates struct {
							Text              string `xml:",chardata"`
							ArrivalDateTime   string `xml:"ArrivalDateTime,attr"`
							DepartureDateTime string `xml:"DepartureDateTime,attr"`
						} `xml:"Dates"`
						FlightSegmentNumbers struct {
							Text                string   `xml:",chardata"`
							FlightSegmentNumber []string `xml:"FlightSegmentNumber"`
						} `xml:"FlightSegmentNumbers"`
					} `xml:"FareComponent"`
					PassengerTypeRequested struct {
						Text string `xml:",chardata"`
						Code string `xml:"Code,attr"`
					} `xml:"PassengerTypeRequested"`
					PassengerTypePriced struct {
						Text string `xml:",chardata"`
						Code string `xml:"Code,attr"`
					} `xml:"PassengerTypePriced"`
				} `xml:"PTC_FareBreakdown"`
				TicketingFees string `xml:"TicketingFees"`
				SignatureLine struct {
					Chardata           string `xml:",chardata"`
					ExpirationDateTime string `xml:"ExpirationDateTime,attr"`
					Text               string `xml:"Text"`
				} `xml:"SignatureLine"`
			} `xml:"AirItineraryPricingInfo"`
		} `xml:"PricedItinerary"`
	} `xml:"ItineraryPricing"`
}

type ReceivedFrom struct {
	Text string `xml:",chardata"`
	Name string `xml:"Name"`
}

type PhoneNumbers struct {
	Text        string `xml:",chardata"`
	PhoneNumber []struct {
		Text      string `xml:",chardata"`
		ElementId string `xml:"elementId,attr"`
		ID        string `xml:"id,attr"`
		Index     string `xml:"index,attr"`
		CityCode  string `xml:"CityCode"`
		Number    string `xml:"Number"`
	} `xml:"PhoneNumber"`
}

type Remarks struct {
	Text   string `xml:",chardata"`
	Remark []struct {
		Text        string `xml:",chardata"`
		ElementId   string `xml:"elementId,attr"`
		ID          string `xml:"id,attr"`
		Index       string `xml:"index,attr"`
		Type        string `xml:"type,attr"`
		RemarkLines struct {
			Text       string `xml:",chardata"`
			RemarkLine struct {
				Chardata string `xml:",chardata"`
				Text     string `xml:"Text"`
			} `xml:"RemarkLine"`
		} `xml:"RemarkLines"`
	} `xml:"Remark"`
}

type GenericSpecialRequests struct {
	Text          string `xml:",chardata"`
	ID            string `xml:"id,attr"`
	MsgType       string `xml:"msgType,attr"`
	Type          string `xml:"type,attr"`
	Code          string `xml:"Code"`
	FreeText      string `xml:"FreeText"`
	ActionCode    string `xml:"ActionCode"`
	NumberInParty string `xml:"NumberInParty"`
	AirlineCode   string `xml:"AirlineCode"`
	FullText      string `xml:"FullText"`
}

type OpenReservationElements struct {
	Text                   string                   `xml:",chardata"`
	OpenReservationElement []OpenReservationElement `xml:"OpenReservationElement"`
}

type OpenReservationElement struct {
	Text           string `xml:",chardata"`
	ElementId      string `xml:"elementId,attr"`
	ID             string `xml:"id,attr"`
	Type           string `xml:"type,attr"`
	ServiceRequest struct {
		Text           string `xml:",chardata"`
		ActionCode     string `xml:"actionCode,attr"`
		AirlineCode    string `xml:"airlineCode,attr"`
		Code           string `xml:"code,attr"`
		ServiceCount   string `xml:"serviceCount,attr"`
		ServiceType    string `xml:"serviceType,attr"`
		SsrType        string `xml:"ssrType,attr"`
		FreeText       string `xml:"FreeText"`
		FullText       string `xml:"FullText"`
		TravelDocument struct {
			Text            string `xml:",chardata"`
			Type            string `xml:"Type"`
			DateOfBirth     string `xml:"DateOfBirth"`
			Gender          string `xml:"Gender"`
			LastName        string `xml:"LastName"`
			FirstName       string `xml:"FirstName"`
			Infant          string `xml:"Infant"`
			HasDocumentData string `xml:"HasDocumentData"`
		} `xml:"TravelDocument"`
	} `xml:"ServiceRequest"`
	Loyalty struct {
		Text          string `xml:",chardata"`
		FrequentFlyer struct {
			Text                 string `xml:",chardata"`
			ActionCode           string `xml:"ActionCode"`
			PreviousActionCode   string `xml:"PreviousActionCode"`
			Vendor               string `xml:"Vendor"`
			ReceivingCarrierCode string `xml:"ReceivingCarrierCode"`
			VitType              string `xml:"VitType"`
			FrequentFlyerNumber  string `xml:"FrequentFlyerNumber"`
			Banner               string `xml:"Banner"`
			Tag                  string `xml:"Tag"`
		} `xml:"FrequentFlyer"`
	} `xml:"Loyalty"`
	NameAssociation struct {
		Text          string `xml:",chardata"`
		LastName      string `xml:"LastName"`
		FirstName     string `xml:"FirstName"`
		NameRefNumber string `xml:"NameRefNumber"`
		ID            string `xml:"Id"`
		ReferenceId   string `xml:"ReferenceId"`
	} `xml:"NameAssociation"`
}
