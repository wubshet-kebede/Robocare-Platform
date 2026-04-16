export const useRoleService = () => {
  const { $api } = useNuxtApp();
  const getRoles = () => {
    return $api("/roles", {
      method: "GET",
    });
  };
  return { getRoles };
};
