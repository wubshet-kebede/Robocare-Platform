export const useInvitationService = () => {
  const { $api } = useNuxtApp();

  const inviteStaff = (data: Record<string, any>) => {
    return $api("/invite-staff", {
      method: "POST",
      body: data,
    });
  };

  //   const signup = (data: Record<string, any>) => {
  //     return $api("/auth/signup", {
  //       method: "POST",
  //       body: data,
  //     });
  //   };
  //   const logout = () => {
  //     return $api("/auth/logout", {
  //       method: "POST",
  //     });
  //   };
  //   const me = () => {
  //     return $api("/me", {
  //       method: "GET",
  //     });
  //   };
  return { inviteStaff };
};
