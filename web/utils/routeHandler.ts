export default function routeTitle(active_route: string): string {
  switch (active_route.trim().toLowerCase()) {
    case "/profile/":
      return "NORA | Profile Requests";
    case "/profile/inbox":
      return "NORA | Incomplete Actions";
    case "/profile/mileage":
      return "NORA | Profile Mileage";
    case "/profile/check_requests":
      return "NORA | Profile Check Requests";
    case "/profile/petty_cash":
      return "NORA | Profile Petty Cash Requests";
    case "/petty_cash/":
      return "NORA | Petty Cash Requests";
    case "/petty_cash/create":
      return "NORA | New Petty Cash";
    case "/petty_cash/detail":
      return "NORA | Petty Cash Detail";
    case "/petty_cash/edit":
      return "NORA | Edit Petty Cash";
    case "/petty_cash/reports/grant":
      return "NORA | Grant Petty Cash";
    case "/petty_cash/reports/user":
      return "NORA | User Petty Cash";
    case "/petty_cash/reports/monthly":
    case "/mileage/":
      return "NORA | Mileage Requests";
    case "/mileage/create":
      return "NORA | New Mileage Request";
    case "/mileage/detail":
      return "NORA | Mileage Detail";
    case "/mileage/edit":
      return "NORA | Edit Mileage";
    case "/mileage/reports/grant":
      return "NORA | Grant Mileage";
    case "/mileage/reports/user":
      return "NORA | User Mileage";
    case "/mileage/reports/monthly":
      return "NORA | Monthly Mileage";
    case "/check_request/":
      return "NORA | Check Requests";
    case "/check_request/create":
      return "NORA | New Check Request";
    case "/check_request/detail":
      return "NORA | Check Request Detail";
    case "/check_request/edit":
      return "NORA | Edit Check Request";
    case "/check_request/reports/grant":
      return "NORA | Grant Check Requests";
    case "/check_request/reports/user":
      return "NORA | User Check Requests";
    case "/check_request/reports/monthly":
      return "NORA | Monthly Check Requests";
    case "/how_to":
      return "NORA | Requests Help";
    case "/users/":
      return "NORA | User Requests";
    default:
      return "NORA | Finance Requests";
  }
}
