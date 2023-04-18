import GrantReportSelect from "@/components/grantReportSelect";
import UnAuthorized from "@/components/unAuthorized";
import { useAppContext } from "@/context/AppContext";
import { Mileage_Overview } from "@/types/mileage";
import { useState } from "react";

export default function GrantMileageRequest() {
  const [requests, setRequests] = useState(new Array<Mileage_Overview>());
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>Mileage Requests by Grant</h1>
      <GrantReportSelect reportType="mileage" setReport={setRequests} />
      <p>{JSON.stringify(requests, null, 2)}</p>
    </main>
  );
}
