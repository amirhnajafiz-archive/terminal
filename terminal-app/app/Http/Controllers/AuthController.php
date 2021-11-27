<?php

namespace App\Http\Controllers;

use App\Jobs\KeepUserUntil;
use App\Models\User;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Redis;
use TheSeer\Tokenizer\Token;

/**
 * Handles the authentication features.
 *
 */
class AuthController extends Controller
{
    /**
     * Method for logging in a user.
     *
     * @param Request $request
     * @return JsonResponse
     */
    public function login(Request $request): JsonResponse
    {
        $name = $request->get('name');

        if (User::query()->where('name', '=', $name)->exists()) {
            $token = (new Token(1, $name, $name));
            Redis::set($name, $token);

            KeepUserUntil::dispatch($name)
                ->delay(now()->addHours(5)); // Keeping the user online for 5 hours

        } else {
            $token = 'None';
        }

        return response()
            ->json([
                'token' => $token
            ]);
    }

    /**
     * Method for logging out a user.
     *
     * @param Request $request
     * @return JsonResponse
     */
    public function logout(Request $request): JsonResponse
    {
        $name = $request->get('name');
        $token = $request->get('token');

        if (Redis::exists($name) && Redis::get($name) == $token)
        {
            Redis::del($name);
        }

        return response()
            ->json([
                'status' => 'logged out'
            ]);
    }
}
