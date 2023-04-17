import PettyCashFrom from "@/components/petty_cash_form";

export default function NewPettyCash() {
  return (
    <main>
      <h1>💵</h1>
      <h1>New Petty Cash Request</h1>
      <PettyCashFrom new_request={true} />
    </main>
  );
}
