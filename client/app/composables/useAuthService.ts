export const useAuthService = () => {
  const { $api } = useNuxtApp();

  const login = (data: Record<string, any>) => {
    return $api("/auth/login", {
      method: "POST",
      body: data,
    });
  };

  const signup = (data: Record<string, any>) => {
    return $api("/auth/signup", {
      method: "POST",
      body: data,
    });
  };
  const logout = () => {
    return $api("/auth/logout", {
      method: "POST",
    });
  };
  return { login, signup, logout };
};
