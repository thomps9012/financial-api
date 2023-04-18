import MonthlyReportSelect from "@/components/monthlyReportSelect";
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
      <h1>Mileage Requests from</h1>
      <MonthlyReportSelect reportType="mileage" setReport={setRequests} />
      <p>{JSON.stringify(requests, null, 2)}</p>
    </main>
  );
}
