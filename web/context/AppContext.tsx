import { Incomplete_Action } from "@/types/actions";
import {
  Axios_Credentials,
  Login_Input,
  User_Public_Info,
} from "@/types/users";
import {
  Dispatch,
  ReactNode,
  SetStateAction,
  createContext,
  useContext,
  useEffect,
  useState,
} from "react";
import axios from "axios";
import { Vehicle } from "@/types/users";
import { Mileage_Overview } from "@/types/mileage";
import { Petty_Cash_Overview } from "@/types/petty_cash";
import { Check_Request_Overview } from "@/types/check_requests";
import { useRouter } from "next/router";
import { setCookie, deleteCookie } from "cookies-next";
import ErrorDisplay from "@/components/errorDisplay";
type Props = {
  children: ReactNode;
};

type App_Context_Type = {
  user_credentials: Axios_Credentials;
  logged_in: boolean;
  user_profile: User_Public_Info;
  incomplete_actions: Incomplete_Action[];
  setActions: Dispatch<SetStateAction<Incomplete_Action[]>>;
  clearAction: (action_id: string) => void;
  login: (login_info: Login_Input) => void;
  logout: () => void;
};

const default_app_context: App_Context_Type = {
  user_credentials: {
    headers: {
      Authorization: "",
    },
    withCredentials: true,
  },
  logged_in: false,
  user_profile: {
    id: "",
    email: "",
    name: "",
    last_login: new Date(),
    vehicles: new Array<Vehicle>(),
    is_active: false,
    admin: false,
    permissions: [""],
    incomplete_actions: new Array<Incomplete_Action>(),
    mileage_requests: new Array<Mileage_Overview>(),
    petty_cash_requests: new Array<Petty_Cash_Overview>(),
    check_requests: new Array<Check_Request_Overview>(),
  },
  incomplete_actions: [],
  setActions: () => {},
  clearAction: () => {},
  login: () => {},
  logout: () => {},
};

const AppContext = createContext<App_Context_Type>(default_app_context);

export function useAppContext() {
  return useContext(AppContext);
}

export function AppProvider({ children }: Props) {
  const router = useRouter();
  useEffect(() => {
    setTimeout(logout_user, 1000 * 60 * 60 * 12);
  }, [router.route]);

  const [user_token, setToken] = useState("");
  const [user_logged_in, setLoggedIn] = useState(false);
  const [user_info, setUserInfo] = useState({
    id: "",
    email: "",
    name: "",
    last_login: new Date(),
    vehicles: new Array<Vehicle>(),
    is_active: false,
    admin: false,
    permissions: [""],
    incomplete_actions: new Array<Incomplete_Action>(),
    mileage_requests: new Array<Mileage_Overview>(),
    petty_cash_requests: new Array<Petty_Cash_Overview>(),
    check_requests: new Array<Check_Request_Overview>(),
  });
  const [incompleteActions, setIncompleteActions] = useState(
    new Array<Incomplete_Action>()
  );
  const clearAction = (action_id: string) => {
    const new_state = incompleteActions.filter(({ id }) => id != action_id);
    setIncompleteActions(new_state);
  };
  const setProfileInfo = async (auth_token: string) => {
    const { data, status, statusText } = await axios.get("/api/me", {
      headers: {
        Authorization: `Bearer ${auth_token}`,
      },
      withCredentials: true,
    });
    if (status != 200 || 201) {
      return <ErrorDisplay message={statusText} error={data} path="GET /me" />;
    }
    setUserInfo(data.data);
    setCookie("auth_credentials", {
      headers: {
        Authorization: `Bearer ${auth_token}`,
      },
      withCredentials: true,
    });
  };
  const login_user = async (user_info: Login_Input) => {
    const { data, status, statusText } = await axios.post(
      "/api/auth/login",
      user_info
    );
    if (status != 200 || 201) {
      return (
        <ErrorDisplay
          message={statusText}
          error={data}
          path="POST /auth/login"
        />
      );
    }
    setLoggedIn(true);
    setToken(data.data.token);
    await setProfileInfo(data.data.token);
  };
  const logout_user = async () => {
    const { data, status, statusText } = await axios.post("/api/auth/logout");
    if (status != 200 || 201) {
      return (
        <ErrorDisplay
          path="POST /auth/logout"
          message={statusText}
          error={data}
        />
      );
    }
    setLoggedIn(false);
    setToken("");
    setUserInfo({
      id: "",
      email: "",
      name: "",
      last_login: new Date(),
      vehicles: new Array<Vehicle>(),
      is_active: false,
      admin: false,
      permissions: [""],
      incomplete_actions: new Array<Incomplete_Action>(),
      mileage_requests: new Array<Mileage_Overview>(),
      petty_cash_requests: new Array<Petty_Cash_Overview>(),
      check_requests: new Array<Check_Request_Overview>(),
    });
    deleteCookie("auth_credentials");
    window.location.assign("/");
  };
  const value = {
    user_profile: user_info,
    user_credentials: {
      headers: {
        Authorization: `Bearer ${user_token}`,
      },
      withCredentials: true,
    },
    logged_in: user_logged_in,
    incomplete_actions: incompleteActions,
    setActions: setIncompleteActions,
    clearAction: clearAction,
    login: login_user,
    logout: logout_user,
  };
  return <AppContext.Provider value={value}>{children}</AppContext.Provider>;
}
