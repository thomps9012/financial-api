import GrantReportSelect from "@/components/grantReportSelect";
import UnAuthorized from "@/components/unAuthorized";
import { useAppContext } from "@/context/AppContext";
import { Petty_Cash_Overview } from "@/types/petty_cash";
import { useState } from "react";

export default function GrantPettyCash() {
  const [requests, setRequests] = useState(new Array<Petty_Cash_Overview>());
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>Petty Cash Requests by Grant</h1>
      <GrantReportSelect reportType="petty_cash" setReport={setRequests} />
      <p>{JSON.stringify(requests, null, 2)}</p>
    </main>
  );
}
