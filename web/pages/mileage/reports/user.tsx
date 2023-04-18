import UnAuthorized from "@/components/unAuthorized";
import UserSelect from "@/components/userSelect";
import { useAppContext } from "@/context/AppContext";
import { Mileage_Overview } from "@/types/mileage";
import { useState } from "react";

export default function MileageUserReport() {
  const [requests, setRequests] = useState(new Array<Mileage_Overview>());
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>User Mileage Requests</h1>
      <UserSelect reportType="mileage" setRequests={setRequests} />
      <p>{JSON.stringify(requests, null, 2)}</p>
    </main>
  );
}
