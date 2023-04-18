import GrantReportSelect from "@/components/grantReportSelect";
import UnAuthorized from "@/components/unAuthorized";
import { useAppContext } from "@/context/AppContext";
import { Check_Request_Overview } from "@/types/check_requests";
import { useState } from "react";

export default function GrantCheckRequest() {
  const [requests, setRequests] = useState(new Array<Check_Request_Overview>());
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>Check Requests by Grant</h1>
      <GrantReportSelect reportType="check" setReport={setRequests} />
      <p>{JSON.stringify(requests, null, 2)}</p>
    </main>
  );
}
