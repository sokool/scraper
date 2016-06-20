//
//	//olx := &crawler.Configuration{
//	//	URL: "http://olx.pl/oferty/",
//	//	Next: "a[class*=pageNextPrev]",
//	//	Object: "a[class*='thumb']",
//	//	Template: crawler.Template{
//	//		"id" : "span[class='nowrap marginright'] span",
//	//		"name" : "div[class*=offerheadinner] h1",
//	//		//"description" : "div[id*=textContent] p",
//	//		"createdAt" : "span[class='pdingleft10 brlefte5']",
//	//		//"price" : crawler.Query(func(p *requestor.Page) string {
//	//		//	o := p.Document().Find("div[class*='pricelabel'] strong").First().Text()
//	//		//	return o
//	//		//}),
//	//	},
//	//}
//	//
//	//
//
//	//immoscout := &crawler.Configuration{
//	//	URL: "http://www.immobilienscout24.de/Suche/S-2/Wohnung-Kauf?enteredFrom=result_list",
//	//	Next: "a[data-is24-qa='paging_bottom_next']",
//	//	Object: "a[class='result-list-entry__brand-title-container']",
//	//	Template: crawler.Template{
//	//		"name" : "#expose-title",
//	//		"price": "div[class*='is24qa-kaufpreis is24-value']",
//	//		"type" : "dd[class*='is24qa-wohnungstyp']",
//	//		"living_space" : "div[class*='is24qa-wohnflaeche-ca is24-value']",
//	//		"rooms": "div[class*='is24qa-zi is24-value']",
//	//		"description" : "pre[class=is24qa-objektbeschreibung]",
//	//	},
//	//}
//
