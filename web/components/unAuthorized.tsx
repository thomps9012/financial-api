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
        textAlign: 'center'
      }}
    >
      <Link href="/" className="archive-btn">
        Back to Safety
      </Link>
      <br />
      <div className="hr-red" />
      <h1 className="reject-btn">
        You are Attempting to Visit an Administrative Section of the Finance
        Request Portal
      </h1>
    </main>
  );
}
