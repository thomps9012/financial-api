import UnAuthorized from "@/components/unAuthorized";
import UserSelect from "@/components/userSelect";
import { useAppContext } from "@/context/AppContext";

export default function CheckRequestUserReport() {
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>User Check Requests</h1>
      <UserSelect reportType="Check" />
    </main>
  );
}
