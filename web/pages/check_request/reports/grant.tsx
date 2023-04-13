import GrantReportSelect from "@/components/grantReportSelect";
import { Check_Request_Overview } from "@/types/check_requests";
import { useState } from "react";

export default function GrantCheckRequest() {
  const [requests, setRequests] = useState(new Array<Check_Request_Overview>());
  return (
    <main>
      <h1>Grant Check Request</h1>
      <GrantReportSelect reportType="Check" setReport={setRequests} />
    </main>
  );
}
