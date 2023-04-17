import MileageForm from "@/components/mileage_form";

export default function NewMileageRequest() {
  return (
    <main>
      <h1 id="slide-in">🚗</h1>
      <h1>New Mileage Request</h1>
      <MileageForm new_request={true} />
    </main>
  );
}
