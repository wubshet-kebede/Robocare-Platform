export const useDepartmentService = () => {
  const { $api } = useNuxtApp();

  const getDepartments = () => {
    return $api("/get-departments", {
      method: "GET",
    });
  };

  return { getDepartments };
};
