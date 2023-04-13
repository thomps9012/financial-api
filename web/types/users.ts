import { Incomplete_Action } from "./actions";
import { Check_Request_Overview } from "./check_requests";
import { Mileage_Overview } from "./mileage";
import { Petty_Cash_Overview } from "./petty_cash";

export interface Vehicle {
  id: string;
  name: string;
  description: string;
}

export interface Vehicle_Input {
  name: string;
  description: string;
}

export interface Login_Input {
  id: string;
  email: string;
  name: string;
}

export interface User {
  id: string;
  email: string;
  name: string;
  last_login: Date;
  vehicles: Vehicle[];
  is_active: boolean;
  admin: boolean;
  permissions: string[];
}

export interface User_Public_Info {
  id: string;
  email: string;
  name: string;
  last_login: Date;
  vehicles: Vehicle[];
  is_active: boolean;
  admin: boolean;
  permissions: string[];
  mileage_requests: Mileage_Overview[];
  petty_cash_requests: Petty_Cash_Overview[];
  check_requests: Check_Request_Overview[];
  incomplete_actions: Incomplete_Action[];
}

export interface User_Name_Info {
  id: string;
  name: string;
  active: boolean;
}

export interface User_Context_Info {
  id: string;
  admin: boolean;
  permissions: string[];
}

export interface Axios_Credentials {
  headers: {
    Authorization: string;
  };
  withCredentials: boolean;
}
