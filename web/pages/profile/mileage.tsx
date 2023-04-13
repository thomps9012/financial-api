import { useAppContext } from "@/context/AppContext";

export default function ProfileMileagePage() {
  const { user_profile } = useAppContext();
  const { id } = user_profile;
  return (
    <main>
      <h1>Mileage Page for {id}</h1>
    </main>
  );
}
