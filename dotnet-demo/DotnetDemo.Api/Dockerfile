FROM --platform=$BUILDPLATFORM mcr.microsoft.com/dotnet/sdk:8.0-jammy AS build
WORKDIR /source

# copy csproj and restore as distinct layers
COPY *.csproj .
RUN dotnet restore

# copy everything else and build app
COPY . .
RUN dotnet publish -o /app


# final stage/image
FROM mcr.microsoft.com/dotnet/aspnet:8.0-jammy-chiseled
EXPOSE 8080
WORKDIR /app
COPY --from=build /app .
ENTRYPOINT ["./DotnetDemo.Api"]