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
      <h1>Monthly Mileage Request</h1>
      <MonthlyReportSelect reportType="Mileage" setReport={setRequests} />
    </main>
  );
}
