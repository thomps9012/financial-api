import MonthlyReportSelect from "@/components/monthlyReportSelect";
import { Mileage_Overview } from "@/types/mileage";
import { useState } from "react";

export default function GrantMileageRequest() {
  const [requests, setRequests] = useState(new Array<Mileage_Overview>());
  return (
    <main>
      <h1>Monthly Mileage Request</h1>
      <MonthlyReportSelect reportType="Mileage" setReport={setRequests} />
    </main>
  );
}
