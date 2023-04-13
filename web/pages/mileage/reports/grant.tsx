import GrantReportSelect from "@/components/grantReportSelect";
import { Mileage_Overview } from "@/types/mileage";
import { useState } from "react";

export default function GrantMileageRequest() {
  const [requests, setRequests] = useState(new Array<Mileage_Overview>());
  return (
    <main>
      <h1>Grant Mileage Request</h1>
      <GrantReportSelect reportType="Mileage" setReport={setRequests} />
    </main>
  );
}
