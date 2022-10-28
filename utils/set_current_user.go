package utils

import "financial-api/models/user"

func UserEmailHandler(category user.Category, current_status user.Status, exec_review bool) string {
	var to_email string
	if exec_review {
		to_email = "abradley@norainc.org"
	} else if current_status == user.SUPERVISOR_APPROVED {
		to_email = "finance_requests@norainc.org"
	} else if current_status == user.MANAGER_APPROVED {
		switch category {
		case user.LORAIN:
			to_email = "jward@norainc.org"
		case user.NEXT_STEP:
			to_email = "cwoods@norainc.org"
		case user.PERKINS:
			to_email = "jward@norainc.org"
		case user.PREVENTION:
			to_email = "cwoods@norainc.org"
		default:
			to_email = "finance_requests@norainc.org"
		}
	} else {
		switch category {
		case user.ADMINISTRATIVE:
			to_email = "bgriffin@norainc.org"
		case user.IOP:
			to_email = "jward@norainc.org"
		case user.INTAKE:
			to_email = "cwoods@norainc.org"
		case user.PEERS:
			to_email = "jward@norainc.org"
		case user.ACT_TEAM:
			to_email = "jjordan@norainc.org"
		case user.IHBT:
			to_email = "bgriffin@norainc.org"
		case user.FINANCE:
			to_email = "lfuentes@norainc.org"
		case user.LORAIN:
			to_email = "rgiusti@norainc.org"
		case user.MENS_HOUSE:
			to_email = "jward@norainc.org"
		case user.NEXT_STEP:
			to_email = "dbaker@norainc.org"
		case user.PERKINS:
			to_email = "churt@norainc.org"
		case user.PREVENTION:
			to_email = "lamanor@norainc.org"
		default:
			to_email = "finance_requests@norainc.org"
		}
	}
	return to_email
}
