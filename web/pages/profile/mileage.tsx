import { useAppContext } from "@/context/AppContext";

export default function ProfileMileagePage() {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Mileage Page for {name}</h1>
    </main>
  );
}
