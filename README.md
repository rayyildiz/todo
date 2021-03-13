# Go & React Template

![ci](https://github.com/rayyildiz/todo/workflows/ci/badge.svg)


Backend: 
---

- [PostreSQL](https://github.com/lib/pq)
- [GoCloud](https://gocloud.dev/) 
- [Github Action](https://github.com/features/actions)

Frontend: 
---

- [Material UI](https://material-ui.com/)
- Typescript 
- React Router Dom
- [Register](web/app/src/Pages/Auth/RegisterPage.tsx), [Login](web/app/src/Pages/Auth/LoginPage.tsx), [Forget Password](web/app/src/Pages/Auth/ForgetPasswordPage.tsx) pages

![""](https://images.rayyildiz.dev/go-react-template.png)

## Configure




```dotenv
DEBUG=true
POSTGRES_CONNECTION=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
```

# Cloud Run

This template is configured for [Google CLoud Build](https://console.cloud.google.com/cloud-build/builds) and ready to deploy to [Cloud Run](https://cloud.google.com/run/).

Useful links:

- <https://cloud.google.com/run/docs/quickstarts/build-and-deploy> 
- <https://medium.com/google-cloud/google-cloud-run-for-go-ec09ddbba111> 
