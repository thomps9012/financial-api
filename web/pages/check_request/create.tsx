import Check_Request_Form from "@/components/check_request_form";

export default function NewCheckRequest() {
  return (
    <main>
      <h1>🗃️</h1>
      <h1>New Check Request</h1>
      <Check_Request_Form new_request={true} />
    </main>
  );
}
