import Link from "next/link";

export default function MileageOverview() {
  return (
    <main>
      <h1>Mileage Requests</h1>
      <Link id="new" href="/mileage/create">
        <p>Create New</p>
      </Link>
    </main>
  );
}
