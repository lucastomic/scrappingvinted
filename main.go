package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "https://www.vinted.es/api/v2/catalog/items?page=1&per_page=96&time=1740249986&search_text=iphone&catalog_ids=&size_ids=&brand_ids=&status_ids=&color_ids=&material_ids="

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	headers := map[string]string{
		"accept":             "application/json, text/plain, */*",
		"accept-language":    "es-fr",
		"cache-control":      "no-cache",
		"pragma":             "no-cache",
		"priority":           "u=1, i",
		"referer":            "https://www.vinted.es/catalog?search_text=iphone&page=1&time=1740249986",
		"sec-ch-ua":          "\"Not(A:Brand\";v=\"99\", \"Google Chrome\";v=\"133\", \"Chromium\";v=\"133\"",
		"sec-ch-ua-mobile":   "?1",
		"sec-ch-ua-platform": "\"Android\"",
		"sec-fetch-dest":     "empty",
		"sec-fetch-mode":     "cors",
		"sec-fetch-site":     "same-origin",
		"user-agent":         "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Mobile Safari/537.36",
		"x-anon-id":          "a625c0be-5d1d-42a5-b395-857b90ab5d4e",
		"x-csrf-token":       "75f6c9fa-dc8e-4e52-a000-e09dd4084b3e",
		"x-money-object":     "true",
		"cookie":             "v_udt=Y3ZYczdjZmc1U3hoNTJrS2ZNWGgrY2pVUXdLVC0tak9OMTJRdUt3NzM2a0hObi0tOTRydi9XeWVNOGE3QUl5WGpUNTM4QT09; anonymous-locale=es-fr; anon_id=a625c0be-5d1d-42a5-b395-857b90ab5d4e; anon_id=a625c0be-5d1d-42a5-b395-857b90ab5d4e; v_sid=d088bb95f7dedd296eb0be79730aa43b; OptanonAlertBoxClosed=2025-02-21T19:28:13.232Z; eupubconsent-v2=CQNLQdgQNLQdgAcABBENBeFsAP_gAAAAAChQJxNX_G__bXlj8X71aftkeY1f99h7rsQxBhbJk-4FyLvW_JwX32EzNAz6pqYKmRIAu3TBIQNlHIDURUCgaogVrSDMaEyUoTNKJ6BkiFMRY2cYCFxvm4lDeQCY5vr991d52R-t7dr83dzyy4hHv3a5_2S1WJCdAYctDfv9bROb-9IOd_x8v4v4_FgAAAAAABAAAAAAAAAAAAAAAAAAABcAAABQSCCAAgABcAFAAVAA4AB4AEEAMgA1AB4AEQAJgAVQA3gB6AD8AISAQwBEgCOAEsAJoAVoAw4BlAGWANkAd8A9gD4gH2AfoBAACKQEXARgAjQBQQCoAFXALmAYoA0QBtADcAHEAQ6AkQBOwChwFHgKRAU2AtgBcgC7wF5gMNAZIAycBlwDOYGsAayA2MBt4DdQHJgOXAeOA9oCEIELwgB0ABwAJABzgEHAJ-Aj0BIoCVgE2gKfAWEAvIBiADFoGQgZGA0YBqYDaAG3AN0AeUA-QB-4EBAIGQQRBBMCDAEKwIXDgGAACIAHAAeABcAEgAPwA0ADnAHcAQCAg4CEAE_AKgAXoA6QCEAEegJFASsAmIBMoCbQFIAKTAV2AtQBdADEAGLAMhAZMA0YBpoDUwGvANoAbYA24Bx8DnQOfAeUA-IB9sD9gP3AgeBBECDAEGwIVjoJQAC4AKAAqABwAEAALoAZABqADwAIgATAAqwBcAF0AMQAbwA9AB-gEMARIAlgBNACjAFaAMMAZQA0QBsgDvAHtAPsA_YCKAIwAUEAq4BYgC5gF5AMUAbQA3ABxADqAIdAReAkQBMgCdgFDgKPgU0BTYCrAFigLYAXAAuQBdoC7wF5gL6AYaAx4BkgDJwGVQMsAy4BnIDVQGsANvAbqA4sByYDlwHjgPaAfWBAECFpAAmAAgANAA5wCxAI9ATaApMBeQDUwG2ANuAc-A8oB8QD9gIHgQYAg2BCshAcAAWABQAFwAVQAuABiADeAHoARwA7wCKAEpAKCAVcAuYBigDaAHUgU0BTYCxQFogLgAXIAycBnIDVQHjgQtJQIgAEAALAAoABwAHgARAAmABVAC4AGKAQwBEgCOAFGAK0AbIA7wB-AFXAMUAdQBDoCLwEiAKPAWKAtgBeYDJwGWAM5AawA28B7QEDyQA8AC4A7gCAAFQAR6AkUBKwCbQFJgMWAbkA8oB-4EEQIMFIGwAC4AKAAqABwAEEAMgA0AB4AEQAJgAVQAxAB-gEMARIAowBWgDKAGiANkAd8A-wD9AIsARgAoIBVwC5gF5AMUAbQA3ACHQEXgJEATsAocBTYCxQFsALgAXIAu0BeYC-gGGgMkAZPAywDLgGcwNYA1kBt4DdQHJgPHAe0BCECFpQBCABcAEgAjgBzgDuAIAASIAsQBrwDtgH_AR6AkUBMQCbQFIAKfAV2AugBeQDFgGTANTAa8A8oB8UD9gP3AgYBA8CCYEGAINgQrAA.f_wAAAAAAAAA;  domain_selected=true; access_token_web=eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImNsaWVudF9pZCI6IndlYiIsImF1ZCI6ImZyLmNvcmUuYXBpIiwiaXNzIjoidmludGVkLWlhbS1zZXJ2aWNlIiwiaWF0IjoxNzQwMjQ5Nzg5LCJzaWQiOiIyNTMzNWM2Yi0xNzQwMTY2MDg5Iiwic2NvcGUiOiJwdWJsaWMiLCJleHAiOjE3NDAyNTY5ODksInB1cnBvc2UiOiJhY2Nlc3MifQ.rQEH8L1y1cD5P0f5bXenZLdQPaqtaJI8aCFCg8eVeyciJKMh5CVIrI7Gx0YgfDOkG5SlL_VDJ2jwxOUWO6wIAASxchDPkMLLDx4E3at9wboZcsY4n-jepBTeha3xNm_4bmUofcXWqZ_6MawpPgj3hlgtmdk2_Z-klAccacooIvrzQuyXRnJknDr5FjbJhdDMMnGX8ZlICvffhPxPmA6d4kq_ZA26HAupy68m1X9pFiSeFoJSu-VXeYMdVnRWxnEzUQbZ_Ymlb6JkjrfKG4TWW8S9ylN06rkBBL8sxrTqvGHE1nH77_0Ocu8TilmClwr8niX18Hc6lLfWrLdV4X-MZg; refresh_token_web=eyJraWQiOiJFNTdZZHJ1SHBsQWp1MmNObzFEb3JIM2oyN0J1NS1zX09QNVB3UGlobjVNIiwiYWxnIjoiUFMyNTYifQ.eyJhcHBfaWQiOjQsImNsaWVudF9pZCI6IndlYiIsImF1ZCI6ImZyLmNvcmUuYXBpIiwiaXNzIjoidmludGVkLWlhbS1zZXJ2aWNlIiwiaWF0IjoxNzQwMjQ5Nzg5LCJzaWQiOiIyNTMzNWM2Yi0xNzQwMTY2MDg5Iiwic2NvcGUiOiJwdWJsaWMiLCJleHAiOjE3NDA4NTQ1ODksInB1cnBvc2UiOiJyZWZyZXNoIn0.YXvbMVqkSjnsYltgcDXN-LIkuQ1s8udptdGYvY8-YOhNM1j77TmtohSUj3vI25sDrWz-NL80tjFDwMZ_T1V9VPzaEXakTIxdLW9coMBi19jMXLUZqBf5qBJQ2cTnY0Xy7GXMX19oSswaP9J4lZm2jwsqOaW4vdx_bmXJvfQHL3T_pbh4MZ3YRCVmpwZxRDXMgjNnT4EFvuBanAMcgk7caC5lMfIeAEL2bMmTOqxd3pFpArAAtiUleiZbiEYc3DFnKsboP3REv5mTBy99-2dgUVSWV1Cm_6BrmPCf_k46LPWPgokQ6-SAFiGrI_JQsZJ2V9N4yZuf4g2ttyJttrnW9Q;  __cf_bm=XPTBZbehFH2xEW.oh.4mF6UdzH2TdmUqMmsFM6B_61g-1740249958-1.0.1.1-n4fQSiSFwVEzIicII30Y7qy1hVhNdt9ACLvnLMfnAZ9hVeWYOOwwH8jLGhztpFRxdarfzw98qj6wuO1rD24nGj2Ie3ZA5rosDZP6F0SnWO8; OptanonConsent=isGpcEnabled=0&datestamp=Sat+Feb+22+2025+19%3A46%3A27+GMT%2B0100+(hora+est%C3%A1ndar+de+Europa+central)&version=202312.1.0&browserGpcFlag=0&isIABGlobal=false&consentId=a625c0be-5d1d-42a5-b395-857b90ab5d4e&interactionCount=7&hosts=&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A1%2CC0003%3A1%2CC0004%3A1%2CC0005%3A1%2CV2STACK42%3A1%2CC0015%3A1%2CC0035%3A1&genVendors=V2%3A1%2CV1%3A1%2C&geolocation=ES%3BMD&AwaitingReconsent=false; _gcl_aw=GCL.1740249988.CjwKCAiAiOa9BhBqEiwABCdG89KwiCzA-IuhJnAkGM5AggrXc6mcGCET-3dO7n7aJaiVCVH5bByqHRoC4bsQAvD_BwE; datadome=1kC_l01PO6LdkgzZH0Sha_kq6rlPyLKn5DV2gUDRkO1lWwly0lfZP76ZpOIlppR2x4lN3rnOczYbr714tLcC5XIWQortkoi6ugdMSyhq6m735nfDhVKTR3wN1ATOWjH~; cto_bundle=1ORPWV9RcEZ5bUVJVnBlbyUyRlklMkJ3SjFtaSUyRmdLbVlXa0hBaDNSOFFDaDhTblN4V3FXZmVBJTJCTXg1VGxEazFVREd2cHcxM2dRemJIZ0lnWFMlMkJXT0dzTyUyRmNaTTI4ejdaa1pDUEhCbW5kek5xVjglMkJrayUyRmhCR0R3QU1Ja0dORTUlMkJGbzdmRlJrT2ZhSlo2N2NuenU3dnIlMkZuJTJGS242QUsxeWZHdyUyQkslMkJGeDNmS21FZ28lMkZnenpLSDRBbFNXOGRzbiUyQnBDcUFXWFVJVDU5aDZVSlBSUmJMd21pZ3gxdVl2WkxvWE9wNFJ4YzdONXZkZWVxS1V3V3BsSyUyRjJJTyUyQjBVVDlUJTJGM3U0WlMlMkJ6TWpPMGw0NHRmQ1RpajBqOVklMkZoNXlVYUVUNWw1S25HVzZORTF3d1pVUSUyQlZ0cyUyQmJsMU4lMkJyNUFqOFRtY2JjJTJGVXN5Ug; _dd_s=rum=2&id=fcedd2fd-1b85-4612-a682-6de3eb9674a1&created=1740249792815&expire=1740251667145; _ga_5FL8V760H5=GS1.1.1740249792.2.1.1740250775.0.0.0; _ga_ZJHK1N3D75=GS1.1.1740249792.2.1.1740250775.0.0.0; cf_clearance=J_uwt1Uq8SpTXKIVvL2Oa9gBqzEnShMVKf5WZuIc2_c-1740250777-1.2.1.1-6eBUsnYagXL5Oa5CdZU8Nrs2aqaIYVA4IMJABMabGUFSLoFIJMy7pCvAOP9qBJaQH81MwtwW65YYPFiUVWxaW4YJmmewYkP1SAvuGyW0ZvTEUJv8kXoC87pA8PHQE_VQles8I_JMdyIPJOS37a2V9ablTuExaiSzLfeQZBcJYm41nszwbQ_nn_U2SxUyJ6frD0QRPraoYbM2oyqrEHGV8YYaKHUPGufR1id.qORwtxGTkfRwPjHyvcDQticBC_70g9yuDwTiAQMfbOYWB7lv_rb2dJSFqpykX8iafYVtT.A; _vinted_fr_session=T3prVVBPeFl3bEZvaEJKcnBIbHE5aXRuV1lBZVUrYWpjR2VjbWlpbFBvY29BcHRSeWtjL2F2Wkd1ekZlS1g0ZWxCcTFYNStlbUVmV3pUK2g4cHh3WE5tVVlVMEtYM0QxV2NLVEZsbnphdS9NbDh0K0VEdmhCNFVNOWwySlRTalIzVS8yT3JYaGtTeUpiZ2FLak5wTWNyWnhqUXhsb3dkM001YlBhNEprTzJrWURhMy9mZHYwMmhDQSthSDBDMWU4c05zWkVxbjkrT0pZVHU1S3JTTm9ZZWhSWGJMT3VnZ0hVS2hZTktva0tibGNEUUd0Wlg4TlhldldFTit1SlZYWi0tOW00MHFWc0J1RUtZSTBNYnNTc1dzQT09--c97e5d72020467d8f172225192572ef8e1598988; banners_ui_state=PENDING'", // Añade más cookies si es necesario
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var body struct {
		Items []struct {
			Title string `json:"title"`
			Price struct {
				Amount string `json:"amount"`
			} `json:"price"`
		} `json:"items"`
	}

	// Parsea el `json:""`
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		panic(err)
	}

	for _, item := range body.Items {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Price: %s\n", item.Price.Amount)
		fmt.Println()
	}
}
