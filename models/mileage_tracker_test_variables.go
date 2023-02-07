package models

var Test_api_key = "AIzaSyAf7mF7egyl3Ip35hN1n9gXP854_u5-Zsk"
var Test_mileage_points = Mileage_Points{
	Starting_Point: Location{
		Latitude:  -35.27801,
		Longitude: 149.12958,
	},
	Destination: Location{
		Latitude:  -35.28473,
		Longitude: 149.12836,
	},
	LocationPoints: []Location{
		{
			Latitude:  -35.27801,
			Longitude: 149.12958,
		},
		{
			Latitude:  -35.28032,
			Longitude: 149.12907,
		},
		{
			Latitude:  -35.28099,
			Longitude: 149.12929,
		},
		{
			Latitude:  -35.28144,
			Longitude: 149.12984,
		},
		{
			Latitude:  -35.28194,
			Longitude: 149.13003,
		},
		{
			Latitude:  -35.28282,
			Longitude: 149.12956,
		},
		{
			Latitude:  -35.28302,
			Longitude: 149.12881,
		},
		{
			Latitude:  -35.28473,
			Longitude: 149.12836,
		},
	},
}
var Test_encoded_path = "-35.27801,149.12958|-35.28032,149.12907|-35.28099,149.12929|-35.28144,149.12984|-35.28194,149.13003|-35.28282,149.12956|-35.28302,149.12881|-35.28473,149.12836"
var Test_snap_api_call = "https://roads.googleapis.com/v1/snapToRoads?path=" + Test_encoded_path + "&interpolate=true&key=" + Test_api_key
var Test_snap_response = `{
	"snappedPoints":
	  [
		{
		  "location":
			{ "latitude": -35.27800489993019, "longitude": 149.129531998742 },
		  "originalIndex": 0,
		  "placeId": "ChIJr_xl0GdNFmsRsUtUbW7qABM",
		},
		{
		  "location":
			{ "latitude": -35.2784195, "longitude": 149.12946589999999 },
		  "placeId": "ChIJr_xl0GdNFmsRsUtUbW7qABM",
		},
		{
		  "location":
			{ "latitude": -35.2784195, "longitude": 149.12946589999999 },
		  "placeId": "ChIJ6aXGemhNFmsRZE_zHqhmrG4",
		},
		{
		  "location":
			{ "latitude": -35.2792731, "longitude": 149.12933809999998 },
		  "placeId": "ChIJ6aXGemhNFmsRZE_zHqhmrG4",
		},
		{
		  "location":
			{ "latitude": -35.2792731, "longitude": 149.12933809999998 },
		  "placeId": "ChIJTcTdZ2hNFmsRXokM4mWCWfk",
		},
		{
		  "location": { "latitude": -35.279557, "longitude": 149.1292973 },
		  "placeId": "ChIJTcTdZ2hNFmsRXokM4mWCWfk",
		},
		{
		  "location": { "latitude": -35.279557, "longitude": 149.1292973 },
		  "placeId": "ChIJiUfNQmhNFmsRSsAI-1m6y1g",
		},
		{
		  "location":
			{ "latitude": -35.279610999999996, "longitude": 149.1292889 },
		  "placeId": "ChIJiUfNQmhNFmsRSsAI-1m6y1g",
		},
		{
		  "location": { "latitude": -35.2796484, "longitude": 149.1292791 },
		  "placeId": "ChIJiUfNQmhNFmsRSsAI-1m6y1g",
		},
		{
		  "location": { "latitude": -35.2796484, "longitude": 149.1292791 },
		  "placeId": "ChIJ_RyFQ2hNFmsRoHJAbW7qABM",
		},
		{
		  "location":
			{ "latitude": -35.279947299999996, "longitude": 149.1291894 },
		  "placeId": "ChIJ_RyFQ2hNFmsRoHJAbW7qABM",
		},
		{
		  "location":
			{ "latitude": -35.279947299999996, "longitude": 149.1291894 },
		  "placeId": "ChIJOyypT2hNFmsRZBtscGL0htw",
		},
		{
		  "location":
			{ "latitude": -35.280323564795005, "longitude": 149.1290903128365 },
		  "originalIndex": 1,
		  "placeId": "ChIJOyypT2hNFmsRZBtscGL0htw",
		},
		{
		  "location":
			{ "latitude": -35.2803426, "longitude": 149.12908529999999 },
		  "placeId": "ChIJOyypT2hNFmsRZBtscGL0htw",
		},
		{
		  "location":
			{ "latitude": -35.2803426, "longitude": 149.12908529999999 },
		  "placeId": "ChIJr8xRTGhNFmsRzMb-rxgjspc",
		},
		{
		  "location":
			{ "latitude": -35.280409899999995, "longitude": 149.1290699 },
		  "placeId": "ChIJr8xRTGhNFmsRzMb-rxgjspc",
		},
		{
		  "location": { "latitude": -35.28048179999999, "longitude": 149.129062 },
		  "placeId": "ChIJr8xRTGhNFmsRzMb-rxgjspc",
		},
		{
		  "location": { "latitude": -35.2805541, "longitude": 149.1290623 },
		  "placeId": "ChIJr8xRTGhNFmsRzMb-rxgjspc",
		},
		{
		  "location": { "latitude": -35.280626, "longitude": 149.1290712 },
		  "placeId": "ChIJr8xRTGhNFmsRzMb-rxgjspc",
		},
		{
		  "location": { "latitude": -35.280626, "longitude": 149.1290712 },
		  "placeId": "ChIJv5r0smlNFmsR5nunau79Fv4",
		},
		{
		  "location":
			{ "latitude": -35.280695099999996, "longitude": 149.12908489999998 },
		  "placeId": "ChIJv5r0smlNFmsR5nunau79Fv4",
		},
		{
		  "location": { "latitude": -35.2807629, "longitude": 149.1291046 },
		  "placeId": "ChIJv5r0smlNFmsR5nunau79Fv4",
		},
		{
		  "location": { "latitude": -35.2808294, "longitude": 149.1291306 },
		  "placeId": "ChIJv5r0smlNFmsR5nunau79Fv4",
		},
		{
		  "location": { "latitude": -35.2809064, "longitude": 149.1291693 },
		  "placeId": "ChIJv5r0smlNFmsR5nunau79Fv4",
		},
		{
		  "location":
			{ "latitude": -35.280968200000004, "longitude": 149.129208 },
		  "placeId": "ChIJv5r0smlNFmsR5nunau79Fv4",
		},
		{
		  "location":
			{ "latitude": -35.28101395754623, "longitude": 149.1292430025548 },
		  "originalIndex": 2,
		  "placeId": "ChIJv5r0smlNFmsR5nunau79Fv4",
		},
		{
		  "location":
			{ "latitude": -35.28103840000001, "longitude": 149.1292617 },
		  "placeId": "ChIJv5r0smlNFmsR5nunau79Fv4",
		},
		{
		  "location": { "latitude": -35.2810936, "longitude": 149.1293121 },
		  "placeId": "ChIJv5r0smlNFmsR5nunau79Fv4",
		},
		{
		  "location": { "latitude": -35.2810979, "longitude": 149.1293176 },
		  "placeId": "ChIJv5r0smlNFmsR5nunau79Fv4",
		},
		{
		  "location": { "latitude": -35.2810979, "longitude": 149.1293176 },
		  "placeId": "ChIJpYMSrmlNFmsRXkCoIkZxgBg",
		},
		{
		  "location":
			{ "latitude": -35.281152399999996, "longitude": 149.1294256 },
		  "placeId": "ChIJpYMSrmlNFmsRXkCoIkZxgBg",
		},
		{
		  "location":
			{ "latitude": -35.281152399999996, "longitude": 149.1294256 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.2811784, "longitude": 149.1294706 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.2812258, "longitude": 149.1295413 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.2812771, "longitude": 149.12960759999999 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.281332, "longitude": 149.1296695 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.2813904, "longitude": 149.12972670000002 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.281451700000005, "longitude": 149.1297788 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.28146506991143, "longitude": 149.12978858234607 },
		  "originalIndex": 3,
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.28151580000001, "longitude": 149.1298257 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.28158259999999, "longitude": 149.129867 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.281666099999995, "longitude": 149.1299091 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.2817377, "longitude": 149.1299379 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.281810899999996, "longitude": 149.1299602 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.281884999999996, "longitude": 149.1299765 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.28194399606459, "longitude": 149.1299842294294 },
		  "originalIndex": 4,
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.281959799999996, "longitude": 149.12998629999998 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.282035199999996, "longitude": 149.1299895 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.2821254, "longitude": 149.1299851 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location":
			{ "latitude": -35.282199999999996, "longitude": 149.1299743 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.2822739, "longitude": 149.1299573 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.2823468, "longitude": 149.129934 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.2824178, "longitude": 149.1299043 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.2824379, "longitude": 149.1298945 },
		  "placeId": "ChIJ601MoWlNFmsR5mvkfPp2ovA",
		},
		{
		  "location": { "latitude": -35.2824379, "longitude": 149.1298945 },
		  "placeId": "ChIJe9LPnWlNFmsR7mJw8mYHwG0",
		},
		{
		  "location":
			{ "latitude": -35.282472999999996, "longitude": 149.1298835 },
		  "placeId": "ChIJe9LPnWlNFmsR7mJw8mYHwG0",
		},
		{
		  "location": { "latitude": -35.2825375, "longitude": 149.1298525 },
		  "placeId": "ChIJe9LPnWlNFmsR7mJw8mYHwG0",
		},
		{
		  "location":
			{ "latitude": -35.28257309999999, "longitude": 149.1298319 },
		  "placeId": "ChIJe9LPnWlNFmsR7mJw8mYHwG0",
		},
		{
		  "location":
			{ "latitude": -35.28257309999999, "longitude": 149.1298319 },
		  "placeId": "ChIJaUpThGlNFmsRMHWxoH7EOsc",
		},
		{
		  "location":
			{ "latitude": -35.282665400000006, "longitude": 149.12974459999998 },
		  "placeId": "ChIJaUpThGlNFmsRMHWxoH7EOsc",
		},
		{
		  "location":
			{ "latitude": -35.28274030000001, "longitude": 149.1296645 },
		  "placeId": "ChIJaUpThGlNFmsRMHWxoH7EOsc",
		},
		{
		  "location":
			{ "latitude": -35.282809799999995, "longitude": 149.12957799999998 },
		  "placeId": "ChIJaUpThGlNFmsRMHWxoH7EOsc",
		},
		{
		  "location":
			{ "latitude": -35.28282136229385, "longitude": 149.12956142620385 },
		  "originalIndex": 5,
		  "placeId": "ChIJaUpThGlNFmsRMHWxoH7EOsc",
		},
		{
		  "location": { "latitude": -35.2828744, "longitude": 149.1294854 },
		  "placeId": "ChIJaUpThGlNFmsRMHWxoH7EOsc",
		},
		{
		  "location":
			{ "latitude": -35.282922299999996, "longitude": 149.1294044 },
		  "placeId": "ChIJaUpThGlNFmsRMHWxoH7EOsc",
		},
		{
		  "location":
			{ "latitude": -35.282931500000004, "longitude": 149.1293876 },
		  "placeId": "ChIJaUpThGlNFmsRMHWxoH7EOsc",
		},
		{
		  "location": { "latitude": -35.2830263, "longitude": 149.1291788 },
		  "placeId": "ChIJaUpThGlNFmsRMHWxoH7EOsc",
		},
		{
		  "location": { "latitude": -35.2830263, "longitude": 149.1291788 },
		  "placeId": "ChIJyd3JiWlNFmsR9RUq2ySTTZQ",
		},
		{
		  "location": { "latitude": -35.283054, "longitude": 149.1290996 },
		  "placeId": "ChIJyd3JiWlNFmsR9RUq2ySTTZQ",
		},
		{
		  "location": { "latitude": -35.2830794, "longitude": 149.1290094 },
		  "placeId": "ChIJyd3JiWlNFmsR9RUq2ySTTZQ",
		},
		{
		  "location": { "latitude": -35.2830794, "longitude": 149.1290094 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location":
			{ "latitude": -35.28313383700836, "longitude": 149.12893500604946 },
		  "originalIndex": 6,
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location":
			{ "latitude": -35.283134499999996, "longitude": 149.12893409999998 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location":
			{ "latitude": -35.283190399999995, "longitude": 149.1288668 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location": { "latitude": -35.2832503, "longitude": 149.1288041 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location": { "latitude": -35.2833133, "longitude": 149.1287463 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location": { "latitude": -35.2833794, "longitude": 149.128694 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location":
			{ "latitude": -35.283448299999996, "longitude": 149.128647 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location": { "latitude": -35.2835199, "longitude": 149.1286054 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location": { "latitude": -35.2835934, "longitude": 149.1285699 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location":
			{ "latitude": -35.283668899999995, "longitude": 149.12854059999998 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location":
			{ "latitude": -35.28372649999999, "longitude": 149.1285237 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location":
			{ "latitude": -35.28386179999999, "longitude": 149.12849319999998 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location": { "latitude": -35.2839978, "longitude": 149.1284682 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location":
			{ "latitude": -35.2840205, "longitude": 149.12846779999998 },
		  "placeId": "ChIJWSb8ImpNFmsRp_4JAxJFE1A",
		},
		{
		  "location":
			{ "latitude": -35.2840205, "longitude": 149.12846779999998 },
		  "placeId": "ChIJZe8tFWpNFmsR4brZ1vldk2E",
		},
		{
		  "location":
			{ "latitude": -35.2840524, "longitude": 149.12845969999998 },
		  "placeId": "ChIJZe8tFWpNFmsR4brZ1vldk2E",
		},
		{
		  "location":
			{ "latitude": -35.284341500000004, "longitude": 149.1284124 },
		  "placeId": "ChIJZe8tFWpNFmsR4brZ1vldk2E",
		},
		{
		  "location": { "latitude": -35.2843875, "longitude": 149.1284034 },
		  "placeId": "ChIJZe8tFWpNFmsR4brZ1vldk2E",
		},
		{
		  "location": { "latitude": -35.2843875, "longitude": 149.1284034 },
		  "placeId": "ChIJVx7Ta2pNFmsRx9OI9CnN5tI",
		},
		{
		  "location": { "latitude": -35.2845916, "longitude": 149.1283726 },
		  "placeId": "ChIJVx7Ta2pNFmsRx9OI9CnN5tI",
		},
		{
		  "location": { "latitude": -35.2845916, "longitude": 149.1283726 },
		  "placeId": "ChIJtWxAZmpNFmsRlbUCkc6VtnA",
		},
		{
		  "location":
			{ "latitude": -35.28459730000001, "longitude": 149.1283703 },
		  "placeId": "ChIJtWxAZmpNFmsRlbUCkc6VtnA",
		},
		{
		  "location":
			{ "latitude": -35.284728747199374, "longitude": 149.12834860726772 },
		  "originalIndex": 7,
		  "placeId": "ChIJtWxAZmpNFmsRlbUCkc6VtnA",
		},
	  ],
  }`
var Test_formatted_snap_res = Snapped_Points_Response{
	Snapped_Points: []Snapped_Point{
		{Location: Location{Latitude: -35.27800489993019, Longitude: 149.129531998742}, PlaceID: ""},
		{Location: Location{Latitude: -35.2784195, Longitude: 149.12946589999999}, PlaceID: ""},
		{Location: Location{Latitude: -35.2784195, Longitude: 149.12946589999999}, PlaceID: ""},
		{Location: Location{Latitude: -35.2792731, Longitude: 149.12933809999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.2792731, Longitude: 149.12933809999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.279557, Longitude: 149.1292973}, PlaceID: ""},
		{Location: Location{Latitude: -35.279557, Longitude: 149.1292973}, PlaceID: ""},
		{Location: Location{Latitude: -35.279610999999996, Longitude: 149.1292889}, PlaceID: ""},
		{Location: Location{Latitude: -35.2796484, Longitude: 149.1292791}, PlaceID: ""},
		{Location: Location{Latitude: -35.279947299999996, Longitude: 149.1291894}, PlaceID: ""},
		{Location: Location{Latitude: -35.279947299999996, Longitude: 149.1291894}, PlaceID: ""},
		{Location: Location{Latitude: -35.280323564795005, Longitude: 149.1290903128365}, PlaceID: ""},
		{Location: Location{Latitude: -35.2803426, Longitude: 149.12908529999999}, PlaceID: ""},
		{Location: Location{Latitude: -35.2803426, Longitude: 149.12908529999999}, PlaceID: ""},
		{Location: Location{Latitude: -35.280409899999995, Longitude: 149.1290699}, PlaceID: ""},
		{Location: Location{Latitude: -35.28048179999999, Longitude: 149.129062}, PlaceID: ""},
		{Location: Location{Latitude: -35.2805541, Longitude: 149.1290623}, PlaceID: ""},
		{Location: Location{Latitude: -35.280626, Longitude: 149.1290712}, PlaceID: ""},
		{Location: Location{Latitude: -35.280626, Longitude: 149.1290712}, PlaceID: ""},
		{Location: Location{Latitude: -35.280695099999996, Longitude: 149.12908489999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.2807629, Longitude: 149.1291046}, PlaceID: ""},
		{Location: Location{Latitude: -35.2808294, Longitude: 149.1291306}, PlaceID: ""},
		{Location: Location{Latitude: -35.2809064, Longitude: 149.1291693}, PlaceID: ""},
		{Location: Location{Latitude: -35.280968200000004, Longitude: 149.129208}, PlaceID: ""},
		{Location: Location{Latitude: -35.28101395754623, Longitude: 149.1292430025548}, PlaceID: ""},
		{Location: Location{Latitude: -35.28103840000001, Longitude: 149.1292617}, PlaceID: ""},
		{Location: Location{Latitude: -35.2810936, Longitude: 149.1293121}, PlaceID: ""},
		{Location: Location{Latitude: -35.2810979, Longitude: 149.1293176}, PlaceID: ""},
		{Location: Location{Latitude: -35.2810979, Longitude: 149.1293176}, PlaceID: ""},
		{Location: Location{Latitude: -35.281152399999996, Longitude: 149.1294256}, PlaceID: ""},
		{Location: Location{Latitude: -35.281152399999996, Longitude: 149.1294256}, PlaceID: ""},
		{Location: Location{Latitude: -35.2811784, Longitude: 149.1294706}, PlaceID: ""},
		{Location: Location{Latitude: -35.2812258, Longitude: 149.1295413}, PlaceID: ""},
		{Location: Location{Latitude: -35.2812771, Longitude: 149.12960759999999}, PlaceID: ""},
		{Location: Location{Latitude: -35.281332, Longitude: 149.1296695}, PlaceID: ""},
		{Location: Location{Latitude: -35.2813904, Longitude: 149.12972670000002}, PlaceID: ""},
		{Location: Location{Latitude: -35.281451700000005, Longitude: 149.1297788}, PlaceID: ""},
		{Location: Location{Latitude: -35.28146506991143, Longitude: 149.12978858234607}, PlaceID: ""},
		{Location: Location{Latitude: -35.28151580000001, Longitude: 149.1298257}, PlaceID: ""},
		{Location: Location{Latitude: -35.28158259999999, Longitude: 149.129867}, PlaceID: ""},
		{Location: Location{Latitude: -35.281666099999995, Longitude: 149.1299091}, PlaceID: ""},
		{Location: Location{Latitude: -35.2817377, Longitude: 149.1299379}, PlaceID: ""},
		{Location: Location{Latitude: -35.281810899999996, Longitude: 149.1299602}, PlaceID: ""},
		{Location: Location{Latitude: -35.281884999999996, Longitude: 149.1299765}, PlaceID: ""},
		{Location: Location{Latitude: -35.28194399606459, Longitude: 149.1299842294294}, PlaceID: ""},
		{Location: Location{Latitude: -35.281959799999996, Longitude: 149.12998629999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.282035199999996, Longitude: 149.1299895}, PlaceID: ""},
		{Location: Location{Latitude: -35.2821254, Longitude: 149.1299851}, PlaceID: ""},
		{Location: Location{Latitude: -35.282199999999996, Longitude: 149.1299743}, PlaceID: ""},
		{Location: Location{Latitude: -35.2822739, Longitude: 149.1299573}, PlaceID: ""},
		{Location: Location{Latitude: -35.2823468, Longitude: 149.129934}, PlaceID: ""},
		{Location: Location{Latitude: -35.2824178, Longitude: 149.1299043}, PlaceID: ""},
		{Location: Location{Latitude: -35.2824379, Longitude: 149.1298945}, PlaceID: ""},
		{Location: Location{Latitude: -35.2824379, Longitude: 149.1298945}, PlaceID: ""},
		{Location: Location{Latitude: -35.282472999999996, Longitude: 149.1298835}, PlaceID: ""},
		{Location: Location{Latitude: -35.2825375, Longitude: 149.1298525}, PlaceID: ""},
		{Location: Location{Latitude: -35.28257309999999, Longitude: 149.1298319}, PlaceID: ""},
		{Location: Location{Latitude: -35.28257309999999, Longitude: 149.1298319}, PlaceID: ""},
		{Location: Location{Latitude: -35.282665400000006, Longitude: 149.12974459999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.28274030000001, Longitude: 149.1296645}, PlaceID: ""},
		{Location: Location{Latitude: -35.282809799999995, Longitude: 149.12957799999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.28282136229385, Longitude: 149.12956142620385}, PlaceID: ""},
		{Location: Location{Latitude: -35.2828744, Longitude: 149.1294854}, PlaceID: ""},
		{Location: Location{Latitude: -35.282922299999996, Longitude: 149.1294044}, PlaceID: ""},
		{Location: Location{Latitude: -35.282931500000004, Longitude: 149.1293876}, PlaceID: ""},
		{Location: Location{Latitude: -35.2830263, Longitude: 149.1291788}, PlaceID: ""},
		{Location: Location{Latitude: -35.2830263, Longitude: 149.1291788}, PlaceID: ""},
		{Location: Location{Latitude: -35.283054, Longitude: 149.1290996}, PlaceID: ""},
		{Location: Location{Latitude: -35.2830794, Longitude: 149.1290094}, PlaceID: ""},
		{Location: Location{Latitude: -35.2830794, Longitude: 149.1290094}, PlaceID: ""},
		{Location: Location{Latitude: -35.28313383700836, Longitude: 149.12893500604946}, PlaceID: ""},
		{Location: Location{Latitude: -35.283134499999996, Longitude: 149.12893409999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.283190399999995, Longitude: 149.1288668}, PlaceID: ""},
		{Location: Location{Latitude: -35.2832503, Longitude: 149.1288041}, PlaceID: ""},
		{Location: Location{Latitude: -35.2833133, Longitude: 149.1287463}, PlaceID: ""},
		{Location: Location{Latitude: -35.2833794, Longitude: 149.128694}, PlaceID: ""},
		{Location: Location{Latitude: -35.283448299999996, Longitude: 149.128647}, PlaceID: ""},
		{Location: Location{Latitude: -35.2835199, Longitude: 149.1286054}, PlaceID: ""},
		{Location: Location{Latitude: -35.2835934, Longitude: 149.1285699}, PlaceID: ""},
		{Location: Location{Latitude: -35.283668899999995, Longitude: 149.12854059999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.28372649999999, Longitude: 149.1285237}, PlaceID: ""},
		{Location: Location{Latitude: -35.28386179999999, Longitude: 149.12849319999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.2839978, Longitude: 149.1284682}, PlaceID: ""},
		{Location: Location{Latitude: -35.2840205, Longitude: 149.12846779999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.2840205, Longitude: 149.12846779999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.2840524, Longitude: 149.12845969999998}, PlaceID: ""},
		{Location: Location{Latitude: -35.284341500000004, Longitude: 149.1284124}, PlaceID: ""},
		{Location: Location{Latitude: -35.2843875, Longitude: 149.1284034}, PlaceID: ""},
		{Location: Location{Latitude: -35.2843875, Longitude: 149.1284034}, PlaceID: ""},
		{Location: Location{Latitude: -35.2845916, Longitude: 149.1283726}, PlaceID: ""},
		{Location: Location{Latitude: -35.2845916, Longitude: 149.1283726}, PlaceID: ""},
		{Location: Location{Latitude: -35.28459730000001, Longitude: 149.1283703}, PlaceID: ""},
		{Location: Location{Latitude: -35.284728747199374, Longitude: 149.12834860726772}, PlaceID: ""},
	},
}
var Test_origin = "-35.27801,149.12958"
var Test_destination = "-35.28473,149.12836"
var Test_matrix_api_call = "https://maps.googleapis.com/maps/api/distancematrix/json?origins=" + Test_origin + "&destinations=" + Test_destination + "&units=imperial&key=" + Test_api_key
var Test_matrix_response = `{
	"destination_addresses" : [ "Tourist Drive 1, Canberra ACT 2601, Australia" ],
	"origin_addresses" : [ "60 Northbourne Ave, Canberra ACT 2601, Australia" ],
	"rows" : [
	   {
		  "elements" : [
			 {
				"distance" : {
				   "text" : "0.5 mi",
				   "value" : 833
				},
				"duration" : {
				   "text" : "2 mins",
				   "value" : 101
				},
				"status" : "OK"
			 }
		  ]
	   }
	],
	"status" : "OK"
 }
 `

var Test_formatted_matrix_res = Matrix_Response{
	Destination_Addresses: []string{"Tourist Drive 1, Canberra ACT 2601, Australia"},
	Origin_Addresses:      []string{"60 Northbourne Ave, Canberra ACT 2601, Australia"},
	Rows: []Matrix_Row{
		{
			Elements: []Matrix_Elements{
				{
					Distance: Matrix_Sub_Element{
						Text:  "0.5 mi",
						Value: 833,
					},
					Duration: Matrix_Sub_Element{
						Text:  "2 mins",
						Value: 101,
					},
					Status: "OK",
				},
			},
		},
	},
	Status: "OK",
}

var Test_low_variance = ResponseCompare{
	Matrix_Distance:   0.5,
	Traveled_Distance: 0.5215077466938861,
	Variance:          LOW,
	Difference:        0,
}
var Test_med_variance = ResponseCompare{
	Matrix_Distance:   0.5,
	Traveled_Distance: 7.07,
	Variance:          MEDIUM,
	Difference:        7,
}
var Test_high_variance = ResponseCompare{
	Matrix_Distance:   0.5,
	Traveled_Distance: 20.02,
	Variance:          HIGH,
	Difference:        20,
}