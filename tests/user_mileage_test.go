package tests

// Test case when the user has no mileage requests:
// Input: user_id = "abc123"
// Expected output: []Mileage_Overview{}, nil

// Test case when the database connection fails:
// Input: user_id = "def456"
// Mock: Return an error when calling database.Use("mileage_requests")
// Expected output: []Mileage_Overview{}, error

// Test case when the filter does not match any records:
// Input: user_id = "ghi789"
// Mock: Return an empty cursor when calling collection.Find(context.TODO(), filter, opts)
// Expected output: []Mileage_Overview{}, nil

// Test case when the function is able to retrieve the user's mileage requests:
// Input: user_id = "jkl012"
// Mock: Return a cursor with a single document when calling collection.Find(context.TODO(), filter, opts). The document should contain values for all fields except for current_user, current_status, and is_active. Set these fields to arbitrary values.
// Mock: Return the user's name "John Doe" when calling FindUserName(current_user_id).
// Expected output: []Mileage_Overview{{
// 	ID: "abc123",
// 	User_ID: "jkl012",
// 	Date: time.Time{},
// 	Reimbursement: 0.0,
// 	Current_Status: "pending",
// 	Current_User: "John Doe",
// 	Is_Active: true,
// 	}}, nil
