import Link from "next/link";

export default function UnAuthorized() {
  return (
    <main
      style={{
        margin: 100,
        padding: 20,
        display: "flex",
        flexDirection: "column",
        flexWrap: "wrap",
      }}
    >
      <h1>Unauthorized</h1>
      <div className="hr-red" />
      <h1>
        You are Attempting to Visit an Administrative Section of the Finance
        Request Portal
      </h1>
      <br />
      <div>
        <Link href="/" className="archive-btn">
          Back to Safety
        </Link>
      </div>
    </main>
  );
}
