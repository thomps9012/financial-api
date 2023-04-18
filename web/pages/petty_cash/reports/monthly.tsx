import MonthlyReportSelect from "@/components/monthlyReportSelect";
import UnAuthorized from "@/components/unAuthorized";
import { useAppContext } from "@/context/AppContext";
import { Petty_Cash_Overview } from "@/types/petty_cash";
import { useState } from "react";

export default function MonthlyPettyCash() {
  const [requests, setRequests] = useState(new Array<Petty_Cash_Overview>());
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>Petty Cash Requests from</h1>
      <MonthlyReportSelect reportType="petty_cash" setReport={setRequests} />
      <p>{JSON.stringify(requests, null, 2)}</p>
    </main>
  );
}
