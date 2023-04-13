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
      <h1>Access Denied</h1>
      <div className="hr-red" />
      <h1>You are Attempting to Visit a Protected Website</h1>
      <br />
      <div>
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
