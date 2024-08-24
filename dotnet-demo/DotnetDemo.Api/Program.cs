using DotnetDemo.Api;
using Prometheus;

var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

var summary = Metrics.CreateSummary("myapp_request_duration_seconds", "Duration of the request.", new SummaryConfiguration
{
    LabelNames = ["op"],
    Objectives = [new QuantileEpsilonPair(0.9, 0.01), new QuantileEpsilonPair(0.99, 0.001)]
});

// Enable the /metrics page to export Prometheus metrics.
app.MapMetrics();

// Create endpoint that returns the status of the application.
// Placeholder for the health check
app.MapGet("/healthz", () => "OK");

var service = new Factorial();

app.MapGet("/{number}", (int number) => { var service = new Factorial(); var result = service.Calculate(number); return result; });

app.Run();