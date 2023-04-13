import ErrorDisplay from "@/components/errorDisplay";

export default function ErrorPage() {
  return (
    <main>
      <ErrorDisplay message="test error" path="/fake_path" />
    </main>
  );
}
