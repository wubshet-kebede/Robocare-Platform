export const useDepartmentService = () => {
  const { $api } = useNuxtApp();

  const getDepartments = () => {
    return $api("/get-departments", {
      method: "GET",
    });
  };
  const createDepartment = (data: Record<string, any>) => {
    return $api("/register-departments", {
      method: "POST",
      body: JSON.stringify(data),
    });
  };

  return { getDepartments, createDepartment };
};
