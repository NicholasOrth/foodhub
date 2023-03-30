import React from "react";

function LogoutButton() {

  function handleLogout() {
    // Delete the cookie here
    document.cookie = "jwt=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";

    // Redirect the user to the login screen

    }

  return <button onClick={handleLogout}>Logout</button>;
}
export default LogoutButton;
export {};
