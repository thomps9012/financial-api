import UserSelect from "@/components/userSelect";

export default function MileageUserReport() {
  return (
    <main>
      <h1>Mileage Request User Report</h1>
      <UserSelect reportType="Mileage" />
    </main>
  );
}