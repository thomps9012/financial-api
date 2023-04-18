import UnAuthorized from "@/components/unAuthorized";
import UserSelect from "@/components/userSelect";
import { useAppContext } from "@/context/AppContext";
import { Petty_Cash_Overview } from "@/types/petty_cash";
import { useState } from "react";

export default function PettyCashUserReport() {
  const [requests, setRequests] = useState(new Array<Petty_Cash_Overview>());
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>User Petty Cash Requests</h1>
      <UserSelect reportType="petty_cash" setRequests={setRequests} />
      <p>{JSON.stringify(requests, null, 2)}</p>
    </main>
  );
}
