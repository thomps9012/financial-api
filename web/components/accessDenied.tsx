import { useAppContext } from "@/context/AppContext";

export default function AccessDenied() {
  const { login } = useAppContext();
  return (
    <div
      style={{
        margin: 100,
        padding: 20,
        display: "flex",
        flexDirection: "column",
        flexWrap: "wrap",
      }}
    >
      <h1 style={{ textAlign: "center" }} className="REJECTED">
        Access Denied
      </h1>
      <div className="hr-red" />
      <h2 style={{ textAlign: "center" }}>
        You are Attempting to Visit a Protected Website
      </h2>
      <div className="hr-red" />
      <br />
      <div style={{ display: "flex", justifyContent: "center" }}>
        <a
          onClick={(e: any) => {
            e.preventDefault();
            const data = {
              id: "109157735191825776845",
              name: "TEST FINANCE",
              email: "test@example.com",
            };
            login(data);
          }}
          className="archive-btn"
        >
          Sign In
        </a>
      </div>
    </div>
  );
}
