using Microsoft.AspNetCore.Mvc;

namespace Bar.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class BarController : ControllerBase
    {

        [HttpGet]
        public IActionResult Get()
        {
           return Ok("Hi, from .net core ");
        }
    }
}
