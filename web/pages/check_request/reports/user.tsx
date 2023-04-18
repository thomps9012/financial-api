import UnAuthorized from "@/components/unAuthorized";
import UserSelect from "@/components/userSelect";
import { useAppContext } from "@/context/AppContext";
import { Check_Request_Overview } from "@/types/check_requests";
import { useState } from "react";

export default function CheckRequestUserReport() {
  const [requests, setRequests] = useState(new Array<Check_Request_Overview>());
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>User Check Requests</h1>
      <UserSelect reportType="check" setRequests={setRequests} />
      <p>{JSON.stringify(requests, null, 2)}</p>
    </main>
  );
}
