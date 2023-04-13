import MonthlyReportSelect from "@/components/monthlyReportSelect";
import { Check_Request_Overview } from "@/types/check_requests";
import { useState } from "react";

export default function MonthlyCheckRequests() {
  const [requests, setRequests] = useState(new Array<Check_Request_Overview>());
  return (
    <main>
      <h1>Monthly Check Request</h1>
      <MonthlyReportSelect reportType="Check" setReport={setRequests} />
    </main>
  );
}
