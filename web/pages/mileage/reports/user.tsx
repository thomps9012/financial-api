import UnAuthorized from "@/components/unAuthorized";
import UserSelect from "@/components/userSelect";
import { useAppContext } from "@/context/AppContext";

export default function MileageUserReport() {
  const { user_profile } = useAppContext();
  if (!user_profile.admin) {
    return <UnAuthorized />;
  }
  return (
    <main>
      <h1>Mileage Request User Report</h1>
      <UserSelect reportType="Mileage" />
    </main>
  );
}
