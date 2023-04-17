import ErrorDisplay from "@/components/errorDisplay";

export default function ErrorPage() {
  return (
    <main>
      <ErrorDisplay
        error={new Error("test error")}
        message="test error"
        path="/fake_path"
      />
    </main>
  );
}
