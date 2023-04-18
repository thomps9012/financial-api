import MonthlyReportSelect from "@/components/monthlyReportSelect";
import { Check_Request_Overview } from "@/types/check_requests";
import { useState } from "react";

export default function MonthlyCheckRequests() {
  const [requests, setRequests] = useState(new Array<Check_Request_Overview>());
  return (
    <main>
      <h1>Check Requests from</h1>
      <MonthlyReportSelect reportType="check" setReport={setRequests} />
      <p>{JSON.stringify(requests, null, 2)}</p>
    </main>
  );
}
