import MonthlyReportSelect from "@/components/monthlyReportSelect";
import { Petty_Cash_Overview } from "@/types/petty_cash";
import { useState } from "react";

export default function MonthlyPettyCash() {
  const [requests, setRequests] = useState(new Array<Petty_Cash_Overview>());
  return (
    <main>
      <h1>Monthly Petty Cash Request</h1>
      <MonthlyReportSelect reportType="Petty Cash" setReport={setRequests} />
    </main>
  );
}
